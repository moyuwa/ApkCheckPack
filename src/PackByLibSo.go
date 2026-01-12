package main

import (
	"archive/zip"
	"fmt"
	"io"
	"path"
	"regexp"
	"strings"
)

// 全局变量用于收集所有加固检测结果
var (
	allPackResults []string
)

// 清除所有加固检测结果
func clearPackResults() {
	// 预分配足够大的切片容量，避免频繁的内存重新分配
	// 容量设置为1000，根据实际情况调整
	allPackResults = make([]string, 0, 1000)
}

func PackByLibSo(apkReader *zip.Reader) bool {
	// 清除之前的检测结果
	clearPackResults()

	//获取加固特征json
	apkmap := GetApkPackMap()
	if apkmap == nil {
		fmt.Printf("获取加固特征异常\n")
		return false
	}

	// 预编译所有正则表达式，避免在循环内重复编译
	compiledRegexMap := make(map[string]*regexp.Regexp)
	// 遍历所有加固特征，收集需要编译的正则表达式
	for _, value := range apkmap {
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

	// 收集APK中的所有文件路径，构建文件映射和列表，用于更高效的查找
	filePathMap := make(map[string]bool)
	filePathList := make([]string, 0, len(apkReader.File))
	for _, file := range apkReader.File {
		filePathMap[file.Name] = true
		filePathList = append(filePathList, file.Name)
	}

	// 遍历所有加固特征，使用更高效的查找方式
	for key, value := range apkmap {
		// 检测Sopath特征（精确匹配）
		for _, s := range value.Sopath {
			if filePathMap[s] {
				result := fmt.Sprintf("    Sopath  %s -> %s", key, s)
				allPackResults = append(allPackResults, result)
			}
		}
		
		// 检测Soname特征（包含匹配）
		for _, s := range value.Soname {
			for _, fileName := range filePathList {
				if strings.Contains(fileName, s) {
					result := fmt.Sprintf("    Soname  %s -> %s", key, fileName)
					allPackResults = append(allPackResults, result)
				}
			}
		}
		
		// 检测Other特征（包含匹配）
		for _, s := range value.Other {
			for _, fileName := range filePathList {
				if strings.Contains(fileName, s) {
					result := fmt.Sprintf("    Other  %s -> %s", key, fileName)
					allPackResults = append(allPackResults, result)
				}
			}
		}
		
		// 检测Soregex特征（正则表达式匹配）
		for _, s := range value.Soregex {
			if re, ok := compiledRegexMap[s]; ok {
				for _, fileName := range filePathList {
					if re.MatchString(fileName) {
						result := fmt.Sprintf("   Soregex  %s -> %s", key, fileName)
						allPackResults = append(allPackResults, result)
					}
				}
			}
		}
		
		// 检测jclass特征（Java字节码匹配）
		if len(value.Jclass) > 0 {
			// 遍历APK中的所有DEX文件
			for _, file := range apkReader.File {
				if path.Ext(file.Name) != ".dex" {
					continue
				}
				
				// 打开DEX文件
				fileReader, err := file.Open()
				if err != nil {
					continue
				}
				
				// 限制读取大小，避免内存占用过高
				maxSize := int64(*ArgMaxSize) * 1024 * 1024
				limitReader := io.LimitReader(fileReader, maxSize)
				
				// 读取DEX文件内容
				dexData, err := io.ReadAll(limitReader)
				fileReader.Close()
				if err != nil {
					continue
				}
				
				// 转换为字符串，方便匹配
				dexStr := string(dexData)
				
				// 检查jclass属性中定义的模式
				for _, jclassPattern := range value.Jclass {
					// 跳过注释行
					if strings.HasPrefix(strings.TrimSpace(jclassPattern), "//") {
						continue
					}
					
					// 检查是否包含jclass模式
					if strings.Contains(dexStr, jclassPattern) {
						result := fmt.Sprintf("    jclass   %s -> %s (匹配文件: %s)", key, jclassPattern, file.Name)
						allPackResults = append(allPackResults, result)
					}
				}
			}
		}
	}

	// 输出结果
	// 根据是否是内嵌APK调整缩进级别
	mainIndent := "  - "
	secondaryIndent := "    - "
	tertiaryIndent := "      - "
	if isEmbeddedAPK {
		mainIndent = "    - "
		secondaryIndent = "      - "
		tertiaryIndent = "        - "
	}
	
	fmt.Println("\n" + mainIndent + "加固特征扫描结果")
	
	if len(allPackResults) > 0 {
		fmt.Println(secondaryIndent + "加固特征")
		for _, result := range allPackResults {
			// 移除行首的空格，然后添加适当的缩进
			trimmedResult := strings.TrimSpace(result)
			fmt.Printf(tertiaryIndent + "%s\n", trimmedResult)
		}
	} else {
		fmt.Println(secondaryIndent + "未发现加固特征")
	}
	
	// fmt.Println("\n===============================================================")

	return true
}