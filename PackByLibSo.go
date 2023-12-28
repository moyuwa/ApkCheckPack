package main

import (
	"fmt"
	"github.com/avast/apkparser"
	"os"
	"regexp"
	"strings"
)

func PackByLibSo(apkpath string) bool {
	//获取加固特征json
	apkmap := GetApkPackMap()
	if apkmap == nil {
		fmt.Printf("获取加固特征异常\n")
		return false
	}
	//解析apk文件
	apkReader, err := apkparser.OpenZip(apkpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	defer apkReader.Close()

	//获取文件路径列表
	for key, value := range apkmap {
		//fmt.Println("厂商:", key)
		for _, s := range value.Sopath {
			//fmt.Println("Sopath:", s)
			for _, file := range apkReader.File {
				if file.Name == s {
					fmt.Printf("Sopath %s->%s\n", key, file.Name)
				}
			}
		}
		for _, s := range value.Soname {
			//fmt.Println("Soname:", s)
			for _, file := range apkReader.File {
				if strings.Contains(file.Name, s) {
					fmt.Printf("Soname %s->%s\n", key, file.Name)
				}
			}
		}
		for _, s := range value.Other {
			//fmt.Println("Other:", s)
			for _, file := range apkReader.File {
				if strings.Contains(file.Name, s) {
					fmt.Printf("Other %s->%s\n", key, file.Name)
				}
			}
		}
		for _, s := range value.Soregex {
			//fmt.Println("Soregex:", s)
			for _, file := range apkReader.File {
				re := regexp.MustCompile(s)
				if re.MatchString(file.Name) {
					fmt.Printf("Soregex %s->%s\n", key, file.Name)
				}
			}
		}
	}

	return true
}
