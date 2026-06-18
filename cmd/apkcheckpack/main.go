package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/bypass/ApkCheckPack/pkg/anti"
	"github.com/bypass/ApkCheckPack/pkg/certificate"
	"github.com/bypass/ApkCheckPack/pkg/hardcode"
	"github.com/bypass/ApkCheckPack/pkg/pack"
	"github.com/bypass/ApkCheckPack/pkg/sdk"
)

// 命令行参数
var (
	argFilePath    = flag.String("f", "", "指定APK文件路径或目录")
	argCheckRoot   = flag.Bool("root", true, "是否检测ROOT特征")
	argCheckEmu    = flag.Bool("emu", true, "是否检测模拟器特征")
	argCheckDebug  = flag.Bool("debug", true, "是否检测反调试特征")
	argCheckProxy  = flag.Bool("proxy", true, "是否检测代理特征")
	argCheckSDK    = flag.Bool("sdk", true, "是否检测第三方SDK")
	argCheckHard   = flag.Bool("hardcode", false, "是否检测硬编码信息")
	argCheckCert   = flag.Bool("cert", true, "是否检测证书信息")
	argMaxSize     = flag.Int64("maxsize", 500, "单个文件最大扫描大小(MB)")
	argRecursive   = flag.Bool("r", true, "是否递归扫描内嵌APK")
)

func printUsage() {
	fmt.Println("APK检测工具 - 用于检测APK文件中的特征")
	fmt.Println("\n用法:")
	fmt.Println("  ApkCheckPack [参数] -f <APK文件路径>")
	fmt.Println("\n检测类型:")
	fmt.Println("  -root      启用ROOT检测 (默认: 开启)")
	fmt.Println("  -emu       启用模拟器检测 (默认: 开启)")
	fmt.Println("  -debug     启用反调试检测 (默认: 开启)")
	fmt.Println("  -proxy     启用代理检测 (默认: 开启)")
	fmt.Println("  -sdk       启用第三方SDK检测 (默认: 开启)")
	fmt.Println("  -hardcode  启用硬编码检测 (默认: 关闭)")
	fmt.Println("  -cert      启用证书检测 (默认: 开启)")
	fmt.Println("\n扫描控制:")
	fmt.Println("  -maxsize   单个文件最大扫描大小(MB) (默认: 500)")
	fmt.Println("  -r         递归扫描内嵌APK (默认: 开启)")
	fmt.Println("\n示例:")
	fmt.Println("  ApkCheckPack -f test.apk")
	fmt.Println("  ApkCheckPack -f test.apk -hardcode")
	fmt.Println("  ApkCheckPack -f ./apks -r=false -maxsize 100")
}

func main() {
	flag.Usage = printUsage
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	flag.Parse()

	if *helpFlag || len(os.Args) < 2 {
		printUsage()
		return
	}

	if *argFilePath == "" {
		fmt.Println("错误: 必须提供APK文件路径或文件夹路径")
		fmt.Println("使用 -help 查看帮助信息")
		return
	}

	fmt.Println("APK检测工具 - 扫描配置:")
	fmt.Printf("- 文件路径: %s\n", *argFilePath)
	fmt.Printf("- 检测类型: ROOT(%v) 模拟器(%v) 反调试(%v) 代理(%v) SDK(%v) 硬编码(%v) 证书(%v)\n",
		*argCheckRoot, *argCheckEmu, *argCheckDebug, *argCheckProxy, *argCheckSDK, *argCheckHard, *argCheckCert)
	fmt.Printf("- 最大文件大小: %d MB\n", *argMaxSize)
	fmt.Printf("- 递归扫描: %v\n", *argRecursive)
	fmt.Println("---------------------------------------------------")

	info, err := os.Stat(*argFilePath)
	if err != nil {
		fmt.Printf("错误: 无效的路径：%s\n", *argFilePath)
		return
	}

	if info.IsDir() {
		if err := scanAPKFolder(*argFilePath); err != nil {
			fmt.Printf("扫描APK文件夹 %s 失败：%v\n", *argFilePath, err)
		}
	} else {
		if err := scanAPKFile(*argFilePath); err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", *argFilePath, err)
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
		if err := scanAPKFile(filePath); err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", filePath, err)
		}
	}
	return nil
}

func scanAPKFile(filePath string) error {
	apkReader, err := zip.OpenReader(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	defer apkReader.Close()

	fmt.Printf("正在扫描APK文件: %s\n", filePath)
	scanAPKData(&apkReader.Reader, false)

	// 如果启用递归扫描，扫描内嵌APK
	if *argRecursive {
		scanEmbeddedAPKs(&apkReader.Reader)
	}
	return nil
}

// 计算适当的文件大小限制
func calculateSizeLimit(fileSize int64) int64 {
	maxLimit := int64(1024 * 1024 * (*argMaxSize))
	if fileSize > 0 && fileSize < maxLimit {
		return fileSize
	}
	return maxLimit
}

// 扫描内嵌的APK文件
func scanEmbeddedAPKs(apkReader *zip.Reader) {
	for _, file := range apkReader.File {
		if path.Ext(file.Name) != ".apk" {
			continue
		}

		fmt.Printf("\n  - [!]内嵌APK文件: %s\n", file.Name)

		fileReader, err := file.Open()
		if err != nil {
			fmt.Printf("    - 打开内嵌APK失败: %v\n", err)
			continue
		}

		fileSize := int64(file.UncompressedSize64)
		sizeLimit := calculateSizeLimit(fileSize)

		var embeddedReader *zip.Reader
		if sizeLimit > int64(*argMaxSize)*1024*1024 {
			// 大文件使用临时文件
			tempFile, err := os.CreateTemp("", "embedded_apk_*.apk")
			if err != nil {
				fmt.Printf("    - 创建临时文件失败: %v\n", err)
				fileReader.Close()
				continue
			}
			tempFileName := tempFile.Name()
			_, _ = io.Copy(tempFile, io.LimitReader(fileReader, sizeLimit))
			tempFile.Close()
			fileReader.Close()

			defer os.Remove(tempFileName)
			tempFile2, err := os.Open(tempFileName)
			if err != nil {
				fmt.Printf("    - 重新打开临时文件失败: %v\n", err)
				continue
			}
			tempInfo, _ := tempFile2.Stat()
			embeddedReader, err = zip.NewReader(tempFile2, tempInfo.Size())
			tempFile2.Close()
			if err != nil {
				fmt.Printf("    - 解析内嵌APK失败: %v\n", err)
				continue
			}
		} else {
			// 小文件直接加载到内存
			data, err := io.ReadAll(io.LimitReader(fileReader, sizeLimit))
			fileReader.Close()
			if err != nil {
				fmt.Printf("    - 读取内嵌APK内容失败: %v\n", err)
				continue
			}
			embeddedReader, err = zip.NewReader(bytes.NewReader(data), int64(len(data)))
			if err != nil {
				fmt.Printf("    - 解析内嵌APK失败: %v\n", err)
				continue
			}
		}

		if fileSize > 0 {
			fmt.Printf("    - 内嵌APK大小: %.2f MB\n", float64(fileSize)/(1024*1024))
		}

		scanAPKData(embeddedReader, true)
	}
}

// 执行APK数据扫描
func scanAPKData(apkReader *zip.Reader, isEmbedded bool) error {
	maxScanBytes := int64(*argMaxSize) * 1024 * 1024

	// 检测加固特征
	pack.Detect(apkReader, isEmbedded, maxScanBytes)

	// 根据参数控制安全检测
	if *argCheckRoot || *argCheckEmu || *argCheckDebug || *argCheckProxy {
		antiOpts := anti.Options{
			CheckRoot:   *argCheckRoot,
			CheckEmu:    *argCheckEmu,
			CheckDebug:  *argCheckDebug,
			CheckProxy:  *argCheckProxy,
			MaxScanSize: maxScanBytes,
		}
		anti.Detect(apkReader, antiOpts, isEmbedded)
	}

	// 执行SDK检测
	if *argCheckSDK {
		sdk.Detect(apkReader, isEmbedded)
	}

	// 根据参数控制硬编码检测
	if *argCheckHard {
		hardcode.Detect(apkReader, maxScanBytes)
	}

	// 根据参数控制证书检测
	if *argCheckCert {
		certificate.Detect(apkReader, isEmbedded)
	}

	fmt.Printf("\n")
	return nil
}
