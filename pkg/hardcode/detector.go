package hardcode

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// HardCodedResult 硬编码检测结果
type HardCodedResult struct {
	FilePath  string
	Category  string
	Pattern   string
	MatchText string
}

// PatternCategory 预编译的正则表达式类别映射
type PatternCategory struct {
	Pattern    *regexp.Regexp
	Category   string
	PatternStr string
}

// Detect 检测APK中的硬编码敏感信息
func Detect(apkReader *zip.Reader, maxScanSize int64) bool {
	totalFiles := len(apkReader.File)
	fmt.Printf("\n\n--- 硬编码扫描开始 ---")
	fmt.Printf("\n待扫描文件数: %d", totalFiles)
	fmt.Printf("\n单个文件最大扫描大小: %d MB", maxScanSize/1024/1024)

	// 预编译所有正则表达式
	var patternCategories []PatternCategory
	for _, category := range HardCodeRules {
		for _, patternStr := range category.Patterns {
			patternCategories = append(patternCategories, PatternCategory{
				Pattern:    regexp.MustCompile(patternStr),
				Category:   category.Name,
				PatternStr: patternStr,
			})
		}
	}

	var allResults []HardCodedResult
	var mu sync.Mutex

	var processed int64
	progressChan := make(chan int, totalFiles)
	resultChan := make(chan []HardCodedResult, 10)

	// 启动进度显示goroutine
	go func() {
		for range progressChan {
			count := atomic.AddInt64(&processed, 1)
			percentage := float64(count) / float64(totalFiles) * 100
			fmt.Printf("\r硬编码扫描进度: %d/%d (%.1f%%)", count, totalFiles, percentage)
		}
	}()

	workerCount := 8
	var wg sync.WaitGroup
	workChan := make(chan *zip.File, totalFiles)

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

				dexData, err := io.ReadAll(io.LimitReader(fileReader, maxScanSize))
				fileReader.Close()
				if err != nil {
					progressChan <- 1
					continue
				}

				results := searchHardCoded(dexData, file.Name, patternCategories)
				if len(results) > 0 {
					resultChan <- results
				}
				progressChan <- 1
			}
		}()
	}

	for _, file := range apkReader.File {
		workChan <- file
	}
	close(workChan)

	// 等待工作完成并关闭通道
	go func() {
		wg.Wait()
		close(resultChan)
		close(progressChan)
	}()

	// 收集所有结果
	for results := range resultChan {
		mu.Lock()
		allResults = append(allResults, results...)
		mu.Unlock()
	}

	fmt.Printf("\n\n")

	allResults = deduplicateResults(allResults)

	if len(allResults) > 0 {
		outputResults(allResults)
	} else {
		fmt.Printf("\n未发现硬编码特征")
	}

	return true
}

// searchHardCoded 在数据中搜索硬编码规则
func searchHardCoded(dexData []byte, filePath string, patternCategories []PatternCategory) []HardCodedResult {
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

// deduplicateResults 结果去重
func deduplicateResults(results []HardCodedResult) []HardCodedResult {
	seen := make(map[string]bool)
	var uniqueResults []HardCodedResult

	for _, result := range results {
		key := result.FilePath + "|" + result.Category + "|" + result.MatchText
		if !seen[key] {
			seen[key] = true
			uniqueResults = append(uniqueResults, result)
		}
	}

	return uniqueResults
}

// outputResults 输出硬编码检测结果
func outputResults(results []HardCodedResult) {
	fmt.Printf("\n--- 硬编码检测结果 ---")

	// 按类别分组
	categoryMap := make(map[string][]HardCodedResult)
	var categories []string

	for _, result := range results {
		if _, exists := categoryMap[result.Category]; !exists {
			categories = append(categories, result.Category)
		}
		categoryMap[result.Category] = append(categoryMap[result.Category], result)
	}

	sort.Strings(categories)

	totalCount := len(results)
	fmt.Printf("\n\n共检测到 %d 个硬编码特征，分布在 %d 个类别中", totalCount, len(categories))

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

		sort.Strings(files)

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
