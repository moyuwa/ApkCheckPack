package main

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"sync"
)

// 定义硬编码规则分类
type HardCodeCategory struct {
	Name     string
	Patterns []string
}

// 按类别定义硬编码规则
var hardCodeCategories = []HardCodeCategory{
	{
		Name: "HAE敏感信息",
		Patterns: []string{
			`(?i)password\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)passwd\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)pwd\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)username\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)token\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)auth\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)pass\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)login\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)email\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)account\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)access_key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)secret_key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)client_secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)client_id\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)api[_-]?key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)api[_-]?secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)access[_-]?token\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)refresh[_-]?token\s*=\s*['"][^'"]{3,}['"]`,
		},
	},
	{
		Name: "私钥和证书",
		Patterns: []string{
			`-----BEGIN DSA PRIVATE KEY-----`,
			`-----BEGIN EC PRIVATE KEY-----`,
			`-----BEGIN PGP PRIVATE KEY BLOCK-----`,
			`-----BEGIN RSA PRIVATE KEY-----`,
		},
	},
	{
		Name: "API密钥和令牌",
		Patterns: []string{
			`[aA][pP][iI]_?[kK][eE][yY].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
			`api_key=[0-9a-zA-Z]+`,
			`key-[0-9a-zA-Z]{32}`,
			`AIza[0-9A-Za-z\\-_]{35}`,
		},
	},
	{
		Name: "OAuth和认证令牌",
		Patterns: []string{
			`access_token=[0-9a-zA-Z]+`,
			`access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}`,
			`ya29\\.[0-9A-Za-z\\-_]+`,
			`eyJhbGciOiJ`,
			`EAACEdEose0cBA[0-9A-Za-z]+`,
		},
	},
	{
		Name: "云平台凭证",
		Patterns: []string{
			`((?:A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16})`, // AWS
			`amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`, // Amazon MWS
			`[hH][eE][rR][oO][kK][uU].{0,20}[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`, // Heroku
		},
	},
	{
		Name: "服务账号凭证",
		Patterns: []string{
			`"service_account\"`,
			`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`,
			`[0-9]+:AA[0-9A-Za-z\\-_]{33}`,
		},
	},
	{
		Name: "支付相关密钥",
		Patterns: []string{
			`rk_live_[0-9a-zA-Z]{24}`,
			`sk_live_[0-9a-z]{32}`,
			`sk_live_[0-9a-zA-Z]{24}`,
			`sq0atp-[0-9A-Za-z\\-_]{22}`,
			`sq0csp-[0-9A-Za-z\\-_]{43}`,
		},
	},
	{
		Name: "平台服务密钥",
		Patterns: []string{
			`[fF][aA][cC][eE][bB][oO][oO][kK].{0,20}['|\"][0-9a-f]{32}['|\"]`, // Facebook
			`[gG][iI][tT][hH][uU][bB].{0,20}['|\"][0-9a-zA-Z]{35,40}['|\"]`, // GitHub
			`https://hooks\\.slack\\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`, // Slack
			`https:\/\/[a-zA-Z0-9]{40}@github\.com`, // GitHub
		},
	},
	{
		Name: "其他敏感信息",
		Patterns: []string{
			`[sS][eE][cC][rR][eE][tT].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
			`[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@.{1,100}[\"'\\s]`,
			`[0-9a-f]{32}-us[0-9]{1,2}`,
			`da2-[a-z0-9]{26}`,
			`SK[0-9a-fA-F]{32}`,
		},
	},
}

// 正则表达式模式与类别映射结构
type PatternCategory struct {
	Pattern  *regexp.Regexp
	Category string
	PatternStr string
}

// 预编译所有正则表达式并构建模式到类别的映射
func buildPatternCategories() []PatternCategory {
	var patternCategories []PatternCategory
	for _, category := range hardCodeCategories {
		for _, patternStr := range category.Patterns {
			patternCategories = append(patternCategories, PatternCategory{
				Pattern:  regexp.MustCompile(patternStr),
				Category: category.Name,
				PatternStr: patternStr,
			})
		}
	}
	return patternCategories
}

// 硬编码检测结果结构
type HardCodedResult struct {
	FilePath    string
	Category    string
	Pattern     string
	MatchText   string
}

// 搜索硬编码的规则
func SearchHardCoded(dexData []byte, filePath string, patternCategories []PatternCategory) []HardCodedResult {
	var results []HardCodedResult
	for _, pc := range patternCategories {
		matches := pc.Pattern.FindAll(dexData, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				results = append(results, HardCodedResult{
					FilePath:  filePath,
					Category:  pc.Category,
					Pattern:   pc.PatternStr,
					MatchText: string(match),
				})
			}
		}
	}
	return results
}

// 直接打开apk暴力搜索
func ScanAPKHardCoded(apkReader *zip.Reader) bool {
	// 构建预编译的正则表达式类别映射
	patternCategories := buildPatternCategories()
	
	// 存储所有检测结果
	var allResults []HardCodedResult
	
	// 结果通道和进度通道
	resultsChan := make(chan []HardCodedResult, 10)
	progressChan := make(chan int, len(apkReader.File))
	
	// 启动进度显示goroutine
	go func() {
		processed := 0
		totalFiles := len(apkReader.File)
		for range progressChan {
			processed++
			fmt.Printf("\r硬编码扫描进度: %d/%d", processed, totalFiles)
		}
	}()

	// 启动工作池
	var wg sync.WaitGroup
	workChan := make(chan *zip.File, len(apkReader.File))
	
	// 工作池大小，根据CPU核心数调整
	workerCount := 4
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range workChan {
				fileReader, err := file.Open()
				if err != nil {
					progressChan <- 1
					continue
				}
				
				// 使用参数控制文件大小限制
				dexData, err := io.ReadAll(io.LimitReader(fileReader, 1024*1024*(*ArgMaxSize)))
				fileReader.Close()
				if err != nil {
					progressChan <- 1
					continue
				}
				
				// 搜索当前文件中的硬编码，并将结果发送到通道
				results := SearchHardCoded(dexData, file.Name, patternCategories)
				resultsChan <- results
				progressChan <- 1
			}
		}()
	}

	// 将文件发送到工作通道
	for _, file := range apkReader.File {
		workChan <- file
	}
	close(workChan)

	// 等待所有工作完成
	go func() {
		wg.Wait()
		close(resultsChan)
		close(progressChan)
	}()

	// 收集结果
	for results := range resultsChan {
		allResults = append(allResults, results...)
	}
	
	// 结果去重
	allResults = deduplicateResults(allResults)
	
	// 输出检测结果
	if len(allResults) > 0 {
		outputHardCodedResultsAsText(allResults)
	} else {
		fmt.Printf("\n\n未发现硬编码特征\n")
	}

	return true
}

// 结果去重函数
func deduplicateResults(results []HardCodedResult) []HardCodedResult {
	seen := make(map[string]bool)
	var uniqueResults []HardCodedResult
	
	for _, result := range results {
		// 使用文件路径、类别和匹配文本的组合作为唯一键
		key := result.FilePath + "|" + result.Category + "|" + result.MatchText
		if !seen[key] {
			seen[key] = true
			uniqueResults = append(uniqueResults, result)
		}
	}
	
	return uniqueResults
}

// 文本格式输出硬编码结果
func outputHardCodedResultsAsText(results []HardCodedResult) {
	fmt.Printf("\n--- 硬编码检测结果 ---")
	
	// 按类别分组显示结果
	categoryMap := make(map[string][]HardCodedResult)
	var categories []string
	
	for _, result := range results {
		if _, exists := categoryMap[result.Category]; !exists {
			categories = append(categories, result.Category)
		}
		categoryMap[result.Category] = append(categoryMap[result.Category], result)
	}
	
	// 按类别名称排序
	sort.Strings(categories)
	
	// 输出总统计信息
	totalCount := len(results)
	fmt.Printf("\n\n共检测到 %d 个硬编码特征，分布在 %d 个类别中", totalCount, len(categories))
	
	// 输出分组结果
	for _, category := range categories {
		catResults := categoryMap[category]
		fmt.Printf("\n\n[%s] - %d 个特征", category, len(catResults))
		
		// 按文件名分组
		fileMap := make(map[string][]HardCodedResult)
		var files []string
		
		for _, result := range catResults {
			if _, exists := fileMap[result.FilePath]; !exists {
				files = append(files, result.FilePath)
			}
			fileMap[result.FilePath] = append(fileMap[result.FilePath], result)
		}
		
		// 按文件名排序
		sort.Strings(files)
		
		// 输出每个文件的结果
		for _, file := range files {
			fileResults := fileMap[file]
			fmt.Printf("\n\n  文件: %s - %d 个特征", file, len(fileResults))
			for i, result := range fileResults {
				fmt.Printf("\n    特征 %d:", i+1)
				fmt.Printf("\n      规则: %s", result.Pattern)
				fmt.Printf("\n      匹配: %s", result.MatchText)
				fmt.Printf("\n      %s", strings.Repeat("-", 60))
			}
		}
	}
	
	fmt.Printf("\n\n--- 硬编码检测结束 ---")
	fmt.Printf("\n\n")
}