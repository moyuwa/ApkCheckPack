package sdk

import (
	"archive/zip"
	"fmt"
	"sort"
	"strings"
)

// SDKSoInfo SDK检测结果
type SDKSoInfo struct {
	Team   string
	Label  string
	Soname string
}

// Detect 检测APK中的第三方SDK
func Detect(apkReader *zip.Reader, isEmbedded bool) bool {
	var sdkResult []SDKSoInfo

	for _, file := range apkReader.File {
		if strings.HasSuffix(strings.ToLower(file.Name), ".so") {
			for _, sdk := range SDKRules {
				if strings.Contains(file.Name, sdk.Soname) {
					sdkResult = append(sdkResult, SDKSoInfo{
						Team:   sdk.GetTeam(),
						Label:  sdk.GetLabel(),
						Soname: file.Name,
					})
				}
			}
		}
	}

	if len(sdkResult) > 0 {
		outputResults(sdkResult, isEmbedded)
		return true
	}
	return false
}

func outputResults(sdkList []SDKSoInfo, isEmbedded bool) {
	mainIndent := "  - "
	if isEmbedded {
		mainIndent = "    - "
	}

	fmt.Printf("\n" + mainIndent + "第三方SDK特征扫描结果")

	teamMap := make(map[string][]string)
	var teams []string

	for _, value := range sdkList {
		team := value.Team
		if _, exists := teamMap[team]; !exists {
			teams = append(teams, team)
		}
		teamMap[team] = append(teamMap[team], fmt.Sprintf("%s -> %s", value.Label, value.Soname))
	}

	sort.Strings(teams)

	secondaryIndent := "    - "
	tertiaryIndent := "      - "
	if isEmbedded {
		secondaryIndent = "      - "
		tertiaryIndent = "        - "
	}

	for _, team := range teams {
		fmt.Printf("\n"+secondaryIndent+"%s", team)
		sort.Strings(teamMap[team])
		for _, item := range teamMap[team] {
			fmt.Printf("\n"+tertiaryIndent+"%s", item)
		}
	}
}
