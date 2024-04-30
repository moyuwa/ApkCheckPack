package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var arg_scanhardcode = flag.Bool("s", false, "-s=true\t是否扫描全文件的硬编码信息")
var arg_filepath = flag.String("f", "", "-f test.apk\t指定apk文件路径或目录")

//var dirpath = flag.String("d", "", "-d D:/apkfiles")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供APK文件路径或文件夹路径作为命令行参数s。")
		fmt.Println("用法：ApkCheckPack.exe -s=false -f test.apk")
		return
	}
	flag.Parse()
	//fmt.Println("命令行参数arg_filepath的值：", *arg_filepath)
	//fmt.Println("命令行参数arg_scanhardcode的值：", *arg_scanhardcode)
	//fmt.Println("参数：", flag.Args())

	info, err := os.Stat(*arg_filepath)
	if err != nil {
		fmt.Printf("无效的路径：%s\n", *arg_filepath)
		return
	}

	if info.IsDir() {
		err := scanAPKFolder(*arg_filepath)
		if err != nil {
			fmt.Printf("扫描APK文件夹 %s 失败：%v\n", *arg_filepath, err)
		}
	} else {
		err := scanAPKFile(*arg_filepath)
		if err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", *arg_filepath, err)
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
	verifyApk(filePath)
	PackByLibSo(filePath)
	ScanAPKAnti(filePath)
	if *arg_scanhardcode {
		ScanAPKHardCoded(filePath) //匹配规则待调整
	}

	return nil
}
