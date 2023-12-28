package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供APK文件路径或文件夹路径作为命令行参数。")
		return
	}

	for _, arg := range os.Args[1:] {
		info, err := os.Stat(arg)
		if err != nil {
			fmt.Printf("无效的路径：%s\n", arg)
			continue
		}

		if info.IsDir() {
			err := scanAPKFolder(arg)
			if err != nil {
				fmt.Printf("扫描APK文件夹 %s 失败：%v\n", arg, err)
			}
		} else {
			err := scanAPKFile(arg)
			if err != nil {
				fmt.Printf("扫描APK文件 %s 失败：%v\n", arg, err)
			}
		}
	}
}

func scanAPKFolder(folderPath string) error {
	fileList := []string{}

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".apk") {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	for _, filePath := range fileList {
		fmt.Printf("扫描APK文件 %s\n", filePath)
		err := scanAPKFile(filePath)
		if err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", filePath, err)
		}
	}

	return nil
}

func scanAPKFile(filePath string) error {
	//fmt.Printf("scanAPKFile")
	PackByLibSo(filePath)
	return nil
}
