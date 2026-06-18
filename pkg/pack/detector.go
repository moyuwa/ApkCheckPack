package pack

import (
	"archive/zip"
	"fmt"
	"io"
	"path"
	"regexp"
	"strings"
)

// Detect 检测APK中的加固特征
func Detect(apkReader *zip.Reader, isEmbedded bool, maxScanSize int64) bool {
	allResults := make([]string, 0, 1000)

	// 预编译所有正则表达式
	compiledRegexMap := make(map[string]*regexp.Regexp)
	for _, value := range PackRules {
		for _, s := range value.Soregex {
			if _, ok := compiledRegexMap[s]; !ok {
				re, err := regexp.Compile(s)
				if err != nil {
					fmt.Printf("编译正则表达式失败: %s -> %v\n", s, err)
					continue
				}
				compiledRegexMap[s] = re
			}
		}
	}

	// 收集APK中的所有文件路径和映射
	filePathMap := make(map[string]bool)
	filePathList := make([]string, 0, len(apkReader.File))
	for _, file := range apkReader.File {
		filePathMap[file.Name] = true
		filePathList = append(filePathList, file.Name)
	}

	// 遍历所有加固规则进行检测
	for name, rule := range PackRules {
		// 检测Sopath特征（精确匹配）
		for _, s := range rule.Sopath {
			if filePathMap[s] {
				result := fmt.Sprintf("    Sopath  %s -> %s", name, s)
				allResults = append(allResults, result)
			}
		}

		// 检测Soname特征（包含匹配）
		for _, s := range rule.Soname {
			for _, fileName := range filePathList {
				if strings.Contains(fileName, s) {
					result := fmt.Sprintf("    Soname  %s -> %s", name, fileName)
					allResults = append(allResults, result)
				}
			}
		}

		// 检测Other特征（包含匹配）
		for _, s := range rule.Other {
			for _, fileName := range filePathList {
				if strings.Contains(fileName, s) {
					result := fmt.Sprintf("    Other   %s -> %s", name, fileName)
					allResults = append(allResults, result)
				}
			}
		}

		// 检测Soregex特征（正则表达式匹配）
		for _, s := range rule.Soregex {
			if re, ok := compiledRegexMap[s]; ok {
				for _, fileName := range filePathList {
					if re.MatchString(fileName) {
						result := fmt.Sprintf("   Soregex  %s -> %s", name, fileName)
						allResults = append(allResults, result)
					}
				}
			}
		}

		// 检测Jclass特征（Java字节码匹配）
		if len(rule.Jclass) > 0 {
			for _, file := range apkReader.File {
				if path.Ext(file.Name) != ".dex" {
					continue
				}

				fileReader, err := file.Open()
				if err != nil {
					continue
				}

				dexData, err := io.ReadAll(io.LimitReader(fileReader, maxScanSize))
				fileReader.Close()
				if err != nil {
					continue
				}

				dexStr := string(dexData)

				for _, jclassPattern := range rule.Jclass {
					// 跳过注释行
					if strings.HasPrefix(strings.TrimSpace(jclassPattern), "//") {
						continue
					}

					if strings.Contains(dexStr, jclassPattern) {
						result := fmt.Sprintf("    jclass   %s -> %s (匹配文件: %s)", name, jclassPattern, file.Name)
						allResults = append(allResults, result)
					}
				}
			}
		}
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

	fmt.Println("\n" + mainIndent + "加固特征扫描结果")

	if len(allResults) > 0 {
		fmt.Println(secondaryIndent + "加固特征")
		for _, result := range allResults {
			trimmedResult := strings.TrimSpace(result)
			fmt.Printf(tertiaryIndent+"%s\n", trimmedResult)
		}
	} else {
		fmt.Println(secondaryIndent + "未发现加固特征")
	}

	return true
}
