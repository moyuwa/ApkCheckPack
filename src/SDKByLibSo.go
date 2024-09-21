package main

import (
	"fmt"
	"github.com/avast/apkparser"
	"os"
	"sort"
	"strings"
)

func SDKByLibSo(apkpath string) bool {
	//获取SDK特征json
	apkmap := GetApkSDKMap()
	if apkmap == nil {
		fmt.Printf("获取第三方SDK特征异常\n")
		return false
	}
	//解析apk文件
	apkReader, err := apkparser.OpenZip(apkpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	defer apkReader.Close()

	// 将特征存储到数组，排序后输出
	var sdksolist []ApkSDKJson
	for _, value := range apkmap {
		for _, file := range apkReader.File {
			if strings.Contains(file.Name, value.Soname) {
				//fmt.Printf("发现SDK特征 Soname [%s]->%s\n", value.Team, file.Name)
				sdksolist = append(sdksolist, value)
			}
		}

	}

	//输出匹配结果 先格式化再输出
	fmt.Printf("\n===扫描第三方SDK特征结果===\n\n")

	var pftstr []string
	for _, value := range sdksolist {
		//fmt.Printf("%s, %s->%s", value.Team, value.Label, value.Soname)
		pftstr = append(pftstr, fmt.Sprintf("%s, %s->%s", value.Team, value.Label, value.Soname))
	}

	sort.Strings(pftstr)
	for _, value := range pftstr {
		fmt.Printf("%s\n", value)
	}
	fmt.Printf("\n=======================\n")

	return true
}
