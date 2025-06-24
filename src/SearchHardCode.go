package main

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
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

// 获取所有硬编码规则
func getAllHardCodedPatterns() []string {
	var patterns []string
	for _, category := range hardCodeCategories {
		patterns = append(patterns, category.Patterns...)
	}
	return patterns
}

// 解析硬编码规则列表为正则表达式
func parseHardCodedPatterns(patterns []string) []*regexp.Regexp {
	var regexps []*regexp.Regexp
	for _, pattern := range patterns {
		regexps = append(regexps, regexp.MustCompile(pattern))
	}
	return regexps
}

// 硬编码检测结果结构
type HardCodedResult struct {
	FilePath    string
	Category    string
	Pattern     string
	MatchText   string
}

// 全局变量用于收集所有硬编码检测结果
var allHardCodedResults []HardCodedResult

// 清除所有硬编码检测结果
func clearHardCodedResults() {
	allHardCodedResults = []HardCodedResult{}
}

// 搜索硬编码的规则
func SearchHardCoded(dexData []byte, filePath string, patterns []*regexp.Regexp) {
	patternList := getAllHardCodedPatterns()
	
	for i, pattern := range patterns {
		matches := pattern.FindAll(dexData, -1)
		if len(matches) > 0 {
			// 查找该pattern属于哪个类别
			category := "未分类"
			for _, cat := range hardCodeCategories {
				for _, catPattern := range cat.Patterns {
					if catPattern == patternList[i] {
						category = cat.Name
						break
					}
				}
				if category != "未分类" {
					break
				}
			}
			
			for _, match := range matches {
				result := HardCodedResult{
					FilePath:  filePath,
					Category:  category,
					Pattern:   patternList[i],
					MatchText: string(match),
				}
				allHardCodedResults = append(allHardCodedResults, result)
			}
		}
	}
}

// 直接打开apk暴力搜索
func ScanAPKHardCoded(apkReader *zip.Reader) bool {
	// 清除之前的检测结果
	clearHardCodedResults()
	
	patterns := parseHardCodedPatterns(getAllHardCodedPatterns())

	// 读取全部文件扫描
	totalFiles := len(apkReader.File)
	for i, file := range apkReader.File {
		// 显示进度
		fmt.Printf("\r硬编码扫描进度: %d/%d", i+1, totalFiles)

		fileReader, err := file.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		// 使用参数控制文件大小限制
		dexData, err := io.ReadAll(io.LimitReader(fileReader, 1024*1024*(*ArgMaxSize)))
		fileReader.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		SearchHardCoded(dexData, file.Name, patterns)
	}
	
	// 输出检测结果
	if len(allHardCodedResults) > 0 {
		outputHardCodedResultsAsText()
	} else {
		fmt.Printf("\n\n未发现硬编码特征\n")
	}

	return true
}

// 文本格式输出硬编码结果
func outputHardCodedResultsAsText() {
	fmt.Printf("\n\n===================== 硬编码检测结果 =====================\n")
	
	// 按类别分组显示结果
	categoryMap := make(map[string][]HardCodedResult)
	var categories []string
	
	for _, result := range allHardCodedResults {
		if _, exists := categoryMap[result.Category]; !exists {
			categories = append(categories, result.Category)
		}
		categoryMap[result.Category] = append(categoryMap[result.Category], result)
	}
	
	// 按类别名称排序
	sort.Strings(categories)
	
	// 输出分组结果
	for _, category := range categories {
		fmt.Printf("\n[%s]\n", category)
		results := categoryMap[category]
		
		// 按文件名分组
		fileMap := make(map[string][]HardCodedResult)
		var files []string
		
		for _, result := range results {
			if _, exists := fileMap[result.FilePath]; !exists {
				files = append(files, result.FilePath)
			}
			fileMap[result.FilePath] = append(fileMap[result.FilePath], result)
		}
		
		// 按文件名排序
		sort.Strings(files)
		
		// 输出每个文件的结果
		for _, file := range files {
			fmt.Printf("\n  文件: %s\n", file)
			for _, result := range fileMap[file] {
				fmt.Printf("    规则: %-30s\n", result.Pattern)
				fmt.Printf("    匹配: %s\n", result.MatchText)
				fmt.Printf("    %s\n", strings.Repeat("-", 60))
			}
		}
	}
	
	fmt.Printf("\n================================================================\n")
}