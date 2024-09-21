package main

import (
	"fmt"
	"github.com/avast/apkparser"
	"github.com/cheggaaa/pb/v3"
	"os"
	"regexp"
)

// 定义需要搜索的硬编码规则列表
var hardCodedPatterns = []string{
	`"service_account\"`,
	`((?:A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16})`,
	`(xox[pboa]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})`,
	`-----BEGIN DSA PRIVATE KEY-----`,
	`-----BEGIN EC PRIVATE KEY-----`,
	`-----BEGIN PGP PRIVATE KEY BLOCK-----`,
	`-----BEGIN RSA PRIVATE KEY-----`,
	`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`,
	`[0-9]+:AA[0-9A-Za-z\\-_]{33}`,
	`[0-9a-f]{32}-us[0-9]{1,2}`,
	`[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@.{1,100}[\"'\\s]`,
	`[aA][pP][iI]_?[kK][eE][yY].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
	`[fF][aA][cC][eE][bB][oO][oO][kK].{0,20}['|\"][0-9a-f]{32}['|\"]`,
	`[gG][iI][tT][hH][uU][bB].{0,20}['|\"][0-9a-zA-Z]{35,40}['|\"]`,
	`[hH][eE][rR][oO][kK][uU].{0,20}[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`,
	`[sS][eE][cC][rR][eE][tT].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
	`access_token=[0-9a-zA-Z]+`,
	`access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}`,
	`AIza[0-9A-Za-z\\-_]{35}`,
	`amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`,
	`api_key=[0-9a-zA-Z]+`,
	`da2-[a-z0-9]{26}`,
	`EAACEdEose0cBA[0-9A-Za-z]+`,
	`eyJhbGciOiJ`,
	`https://hooks\\.slack\\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	`https:\/\/[a-zA-Z0-9]{40}@github\.com`,
	`key-[0-9a-zA-Z]{32}`,
	`rk_live_[0-9a-zA-Z]{24}`,
	`SK[0-9a-fA-F]{32}`,
	`sk_live_[0-9a-z]{32}`,
	`sk_live_[0-9a-zA-Z]{24}`,
	`sq0atp-[0-9A-Za-z\\-_]{22}`,
	`sq0csp-[0-9A-Za-z\\-_]{43}`,
	`ya29\\.[0-9A-Za-z\\-_]+`,
}

// 解析硬编码规则列表为正则表达式
func parseHardCodedPatterns(patterns []string) []*regexp.Regexp {
	var regexps []*regexp.Regexp
	for _, pattern := range patterns {
		regexps = append(regexps, regexp.MustCompile(pattern))
	}
	return regexps
}

// 搜索硬编码的规则
func SearchHardCoded(dexData []byte, filePath string, patterns []*regexp.Regexp) {
	for _, pattern := range patterns {
		matches := pattern.FindAll(dexData, -1)
		if len(matches) > 0 {
			fmt.Println("发现硬编码信息:", filePath)
			for _, match := range matches {
				fmt.Println("规则:", pattern, "匹配值:", string(match))
			}
			fmt.Println("----------------------------------------------------")
		}
	}
}

// 直接打开apk暴力搜索
func ScanAPKHardCoded(apkpath string) bool {
	//解析apk文件
	apkReader, err := apkparser.OpenZip(apkpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	defer apkReader.Close()

	patterns := parseHardCodedPatterns(hardCodedPatterns)

	// 创建进度条并开始
	bar := pb.StartNew(len(apkReader.File))
	// 读取全部文件扫描
	for _, file := range apkReader.File {
		//fmt.Printf("Scan %s\n", file.Name)
		bar.Increment()
		var dexData = []byte{}
		dexData, err = file.ReadAll(1024 * 1024 * 100) //单个文件读取最大设置为100MB
		SearchHardCoded(dexData, file.Name, patterns)
	}
	// 结束进度条
	bar.Finish()

	return true
}
