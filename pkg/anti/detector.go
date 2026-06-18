package anti

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"
)

// Options 安全检测选项
type Options struct {
	CheckRoot   bool
	CheckEmu    bool
	CheckDebug  bool
	CheckProxy  bool
	MaxScanSize int64
}

// Detect 检测APK中的安全特征
func Detect(apkReader *zip.Reader, opts Options, isEmbedded bool) bool {
	allRootResults := make([]string, 0, 500)
	allEmuResults := make([]string, 0, 500)
	allDebugResults := make([]string, 0, 500)
	allProxyResults := make([]string, 0, 500)

	// 读取dex文件扫描
	for _, file := range apkReader.File {
		if path.Ext(file.Name) != ".dex" {
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 分块读取和处理文件
		scanFileByChunk(fileReader, file.Name, opts,
			&allRootResults, &allEmuResults, &allDebugResults, &allProxyResults)
		fileReader.Close()
	}

	// 输出结果
	mainIndent := "  - "
	secondaryIndent := "    - "
	tertiaryIndent := "      - "
	if isEmbedded {
		mainIndent = "    - "
		secondaryIndent = "      - "
		tertiaryIndent = "        - "
	}

	fmt.Println("\n" + mainIndent + "安全检测特征扫描结果")

	// 统计所有检测结果
	totalResults := 0
	if opts.CheckRoot {
		totalResults += len(allRootResults)
	}
	if opts.CheckEmu {
		totalResults += len(allEmuResults)
	}
	if opts.CheckDebug {
		totalResults += len(allDebugResults)
	}
	if opts.CheckProxy {
		totalResults += len(allProxyResults)
	}

	if totalResults == 0 {
		fmt.Println(secondaryIndent + "未发现安全检测特征")
	} else {
		if opts.CheckRoot && len(allRootResults) > 0 {
			fmt.Println(secondaryIndent + "ROOT检测特征")
			for _, result := range allRootResults {
				trimmedResult := strings.TrimSpace(result)
				fmt.Printf(tertiaryIndent+"%s\n", trimmedResult)
			}
		}

		if opts.CheckEmu && len(allEmuResults) > 0 {
			fmt.Println(secondaryIndent + "模拟器检测特征")
			for _, result := range allEmuResults {
				trimmedResult := strings.TrimSpace(result)
				fmt.Printf(tertiaryIndent+"%s\n", trimmedResult)
			}
		}

		if opts.CheckDebug && len(allDebugResults) > 0 {
			fmt.Println(secondaryIndent + "反调试检测特征")
			for _, result := range allDebugResults {
				trimmedResult := strings.TrimSpace(result)
				fmt.Printf(tertiaryIndent+"%s\n", trimmedResult)
			}
		}

		if opts.CheckProxy && len(allProxyResults) > 0 {
			fmt.Println(secondaryIndent + "代理检测特征")
			for _, result := range allProxyResults {
				trimmedResult := strings.TrimSpace(result)
				fmt.Printf(tertiaryIndent+"%s\n", trimmedResult)
			}
		}
	}

	return true
}

// getMaxPatternLength 计算最长特征长度
func getMaxPatternLength(opts Options) int {
	maxLen := 0

	if opts.CheckRoot {
		for _, pattern := range RootFilePatterns {
			if len(pattern.Pattern) > maxLen {
				maxLen = len(pattern.Pattern)
			}
		}
		for _, pattern := range RootAppPatterns {
			if len(pattern.Pattern) > maxLen {
				maxLen = len(pattern.Pattern)
			}
		}
	}

	if opts.CheckEmu {
		for _, pattern := range EmulatorPatterns {
			if len(pattern.Pattern) > maxLen {
				maxLen = len(pattern.Pattern)
			}
		}
	}

	if opts.CheckDebug {
		for _, pattern := range DebugPatterns {
			if len(pattern.Pattern) > maxLen {
				maxLen = len(pattern.Pattern)
			}
		}
	}

	if opts.CheckProxy {
		for _, pattern := range ProxyPatterns {
			if len(pattern.Pattern) > maxLen {
				maxLen = len(pattern.Pattern)
			}
		}
	}

	return maxLen
}

// scanFileByChunk 分块读取并检测文件内容
func scanFileByChunk(fileReader io.Reader, filePath string, opts Options,
	allRootResults, allEmuResults, allDebugResults, allProxyResults *[]string) {
	maxPatternLen := getMaxPatternLength(opts)

	chunkSize := int64(8 * 1024 * 1024)
	overlapSize := int64(maxPatternLen)

	buffer := make([]byte, chunkSize+overlapSize)
	prevBuffer := make([]byte, overlapSize)

	limitReader := io.LimitReader(fileReader, opts.MaxScanSize)

	for {
		n, err := limitReader.Read(buffer[overlapSize:])
		if n == 0 {
			break
		}

		// 将上一次的末尾部分复制到当前缓冲区开头
		copy(buffer, prevBuffer)

		// 检测当前块中的特征
		scanDexAnti(buffer[:overlapSize+int64(n)], filePath, opts,
			allRootResults, allEmuResults, allDebugResults, allProxyResults)

		// 保存当前块的末尾部分用于下一次检测
		if int64(n) > overlapSize {
			copy(prevBuffer, buffer[int64(n):int64(n)+overlapSize])
		} else {
			copy(prevBuffer, buffer[:int64(n)])
			if int64(n) < overlapSize {
				for i := int64(n); i < overlapSize; i++ {
					prevBuffer[i] = 0
				}
			}
		}

		if err != nil {
			break
		}
	}
}

// scanDexAnti 检测单个数据块中的特征
func scanDexAnti(dexData []byte, filePath string, opts Options,
	allRootResults, allEmuResults, allDebugResults, allProxyResults *[]string) {
	if opts.CheckRoot {
		for _, pattern := range RootFilePatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", filePath, pattern.Pattern, pattern.Description)
				*allRootResults = append(*allRootResults, result)
			}
		}
		for _, pattern := range RootAppPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", filePath, pattern.Pattern, pattern.Description)
				*allRootResults = append(*allRootResults, result)
			}
		}
	}

	if opts.CheckEmu {
		for _, pattern := range EmulatorPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", filePath, pattern.Pattern, pattern.Description)
				*allEmuResults = append(*allEmuResults, result)
			}
		}
	}

	if opts.CheckDebug {
		for _, pattern := range DebugPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", filePath, pattern.Pattern, pattern.Description)
				*allDebugResults = append(*allDebugResults, result)
			}
		}
	}

	if opts.CheckProxy {
		for _, pattern := range ProxyPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", filePath, pattern.Pattern, pattern.Description)
				*allProxyResults = append(*allProxyResults, result)
			}
		}
	}
}
