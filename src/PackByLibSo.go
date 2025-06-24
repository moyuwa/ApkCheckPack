package main

import (
	"archive/zip"
	"fmt"
	"regexp"
	"strings"
)

// 全局变量用于收集所有加固检测结果
var (
	allPackResults []string
)

// 清除所有加固检测结果
func clearPackResults() {
	allPackResults = []string{}
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

	//获取文件路径列表
	for key, value := range apkmap {
		for _, s := range value.Sopath {
			for _, file := range apkReader.File {
				if file.Name == s {
					// 输出并收集结果
					//fmt.Printf("发现加固特征 Sopath  %s->%s\n", key, file.Name)
					result := fmt.Sprintf("    Sopath  %s -> %s", key, file.Name)
					allPackResults = append(allPackResults, result)
				}
			}
		}
		for _, s := range value.Soname {
			for _, file := range apkReader.File {
				if strings.Contains(file.Name, s) {
					// 输出并收集结果
					//fmt.Printf("发现加固特征 Soname  %s->%s\n", key, file.Name)
					result := fmt.Sprintf("    Soname  %s -> %s", key, file.Name)
					allPackResults = append(allPackResults, result)
				}
			}
		}
		for _, s := range value.Other {
			for _, file := range apkReader.File {
				if strings.Contains(file.Name, s) {
					// 输出并收集结果
					//fmt.Printf("发现加固特征 Other   %s->%s\n", key, file.Name)
					result := fmt.Sprintf("    Other  %s -> %s", key, file.Name)
					allPackResults = append(allPackResults, result)
				}
			}
		}
		for _, s := range value.Soregex {
			for _, file := range apkReader.File {
				re := regexp.MustCompile(s)
				if re.MatchString(file.Name) {
					// 输出并收集结果
					//fmt.Printf("发现加固特征 Soregex %s->%s\n", key, file.Name)
					result := fmt.Sprintf("   Soregex  %s -> %s", key, file.Name)
					allPackResults = append(allPackResults, result)
				}
			}
		}
	}

	// 输出结果
	fmt.Println("\n===================== 加固特征扫描结果 =====================")
	
	if len(allPackResults) > 0 {
		fmt.Println("\n[加固特征]")
		for _, result := range allPackResults {
			fmt.Println(result)
		}
	} else {
		fmt.Println("\n未发现加固特征")
	}
	
	// fmt.Println("\n================================================================")

	return true
}