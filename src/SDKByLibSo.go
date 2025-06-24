package main

import (
	"archive/zip"
	"fmt"
	"path"
	"sort"
	"strings"
)

type SDKSoInfo struct {
	Team   string
	Label  string
	Soname string
}

var SDKList struct {
	Team   string
	Label  string
	Soname string
}

var (
	checkSDK = true
	quiet    = false
)

func SDKByLibSo(apkReader *zip.Reader) bool {
	if !checkSDK {
		return false
	}

	//获取SDK特征json
	apksdkjsons := GetApkSDKMap()
	if apksdkjsons == nil {
		fmt.Printf("获取第三方SDK特征异常\n")
		return false
	}

	var sdksolist []SDKSoInfo
	
	for _, file := range apkReader.File {
		if path.Ext(file.Name) == ".so" {
			for _, sdk := range apksdkjsons {
				if strings.Contains(file.Name, sdk.Soname) {
					sdksolist = append(sdksolist, SDKSoInfo{
						Team:   sdk.Team,
						Label:  sdk.Label, 
						Soname: file.Name, // sdk.Soname
					})
				}
			}
		}
	}

	if len(sdksolist) > 0 {
		outputSDKResults(sdksolist)
		return true
	}
	return false
}

func outputSDKResults(sdksolist []SDKSoInfo) {
	if !quiet {
		fmt.Printf("\n===================== 第三方SDK特征扫描结果 =====================\n")
	}


	teamMap := make(map[string][]string)
	var teams []string
	
	for _, value := range sdksolist {
		team := value.Team
		if _, exists := teamMap[team]; !exists {
			teams = append(teams, team)
		}
		teamMap[team] = append(teamMap[team], fmt.Sprintf("    %s -> %s", value.Label, value.Soname))
	}
	
	sort.Strings(teams)
	
	for _, team := range teams {
		fmt.Printf("\n[%s]\n", team)
		sort.Strings(teamMap[team])
		for _, item := range teamMap[team] {
			fmt.Println(item)
		}
	}

	if !quiet {
		// fmt.Printf("\n================================================================\n")
	}
}