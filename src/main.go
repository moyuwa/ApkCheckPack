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
)

// 命令行参数 - 导出为其他包使用
var (
	// 文件路径参数
	ArgFilePath = flag.String("f", "", "指定APK文件路径或目录")
	
	// 检测类型参数
	ArgCheckRoot     = flag.Bool("root", true, "是否检测ROOT特征")
	ArgCheckEmu      = flag.Bool("emu", true, "是否检测模拟器特征")
	ArgCheckDebug    = flag.Bool("debug", true, "是否检测反调试特征")
	ArgCheckProxy    = flag.Bool("proxy", true, "是否检测代理特征")
	ArgCheckSDK      = flag.Bool("sdk", true, "是否检测第三方SDK")
	ArgCheckHardcode = flag.Bool("hardcode", false, "是否检测硬编码信息")
	ArgCheckCert     = flag.Bool("cert", true, "是否检测证书信息")
	
	// 扫描控制参数
	ArgMaxSize      = flag.Int64("maxsize", 500, "单个文件最大扫描大小(MB)")
	ArgRecursive    = flag.Bool("r", true, "是否递归扫描内嵌APK")
)

// 打印使用说明
func printUsage() {
	fmt.Println("APK检测工具 - 用于检测APK文件中的特征")
	fmt.Println("\n用法:")
	fmt.Println("  ApkCheckPack.exe [参数] -f <APK文件路径>")
	
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
	fmt.Println("  ApkCheckPack.exe -f test.apk")
	fmt.Println("  ApkCheckPack.exe -f test.apk -hardcode")
	fmt.Println("  ApkCheckPack.exe -f ./apks -r=false -maxsize 100")
}

func main() {
	// 添加帮助标志
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	flag.Parse()
	
	// 显示帮助信息
	if *helpFlag || len(os.Args) < 2 {
		printUsage()
		return
	}
	
	// 验证必要参数
	if *ArgFilePath == "" {
		fmt.Println("错误: 必须提供APK文件路径或文件夹路径")
		fmt.Println("使用 -help 查看帮助信息")
		return
	}
	
	// 显示扫描配置
	fmt.Println("APK检测工具 - 扫描配置:")
	fmt.Printf("- 文件路径: %s\n", *ArgFilePath)
	fmt.Printf("- 检测类型: ROOT(%v) 模拟器(%v) 反调试(%v) 代理(%v) SDK(%v) 硬编码(%v) 证书(%v)\n", 
		*ArgCheckRoot, *ArgCheckEmu, *ArgCheckDebug, *ArgCheckProxy, *ArgCheckSDK, *ArgCheckHardcode, *ArgCheckCert)
	fmt.Printf("- 最大文件大小: %d MB\n", *ArgMaxSize)
	fmt.Printf("- 递归扫描: %v\n", *ArgRecursive)
	fmt.Println("---------------------------------------------------")
	
	// 检查路径是否有效
	info, err := os.Stat(*ArgFilePath)
	if err != nil {
		fmt.Printf("错误: 无效的路径：%s\n", *ArgFilePath)
		return
	}

	if info.IsDir() {
		err := scanAPKFolder(*ArgFilePath)
		if err != nil {
			fmt.Printf("扫描APK文件夹 %s 失败：%v\n", *ArgFilePath, err)
		}
	} else {
		err := scanAPKFile(*ArgFilePath)
		if err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", *ArgFilePath, err)
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

		//fmt.Printf("扫描APK文件 %s\n", filePath)
		err := scanAPKFile(filePath)
		if err != nil {
			fmt.Printf("扫描APK文件 %s 失败：%v\n", filePath, err)
		}
	}

	return nil
}

func scanAPKFile(filePath string) error {

	// 打开APK文件
	apkReader, err := zip.OpenReader(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	defer apkReader.Close()

	// 扫描主APK
	fmt.Printf("正在扫描APK文件: %s\n", filePath)

	ScanAPKData(&apkReader.Reader)

	// 如果启用递归扫描，扫描内嵌APK
	if *ArgRecursive {
		scanEmbeddedAPKs(&apkReader.Reader)
	}
	
	return nil
}

// 扫描内嵌的APK文件
func scanEmbeddedAPKs(apkReader *zip.Reader) {
	for _, file := range apkReader.File {
		if path.Ext(file.Name) != ".apk" {
			continue
		}

		fmt.Printf("\t发现内嵌APK文件: %s\n", file.Name)
		
		if err := scanSingleEmbeddedAPK(file); err != nil {
			fmt.Printf("\t处理内嵌APK文件 %s 失败: %v\n", file.Name, err)
		}
	}
}

// 计算适当的文件大小限制
func calculateSizeLimit(fileSize int64) int64 {
	// 获取用户设置的最大限制（MB转换为字节）
	maxLimit := int64(1024 * 1024 * (*ArgMaxSize))
	
	// 如果文件大小已知且小于最大限制，使用实际文件大小
	if fileSize > 0 && fileSize < maxLimit {
		return fileSize
	}
	
	// 否则使用最大限制
	return maxLimit
}

// 扫描单个内嵌APK文件
func scanSingleEmbeddedAPK(file *zip.File) error {
	// 获取文件大小
	fileSize := int64(file.UncompressedSize64)
	
	// 打开内嵌APK文件
	fileReader, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开内嵌APK失败: %v", err)
	}
	
	// 使用匿名函数确保资源正确释放
	err = func() error {
		defer fileReader.Close()
		
		// 计算适当的大小限制
		sizeLimit := calculateSizeLimit(fileSize)
		
		// 如果文件大小已知，打印信息
		if fileSize > 0 {
			fmt.Printf("\t内嵌APK大小: %.2f MB, 读取限制: %.2f MB\n", 
				float64(fileSize)/(1024*1024), 
				float64(sizeLimit)/(1024*1024))
		}
		
		// 创建限制大小的读取器
		limitReader := io.LimitReader(fileReader, sizeLimit)
		
		// 分块读取APK内容
		var buffer bytes.Buffer
		_, err := io.Copy(&buffer, limitReader)
		if err != nil {
			return fmt.Errorf("读取内嵌APK内容失败: %v", err)
		}
		
		// 创建zip reader
		apkData := buffer.Bytes()
		embeddedReader, err := zip.NewReader(bytes.NewReader(apkData), int64(len(apkData)))
		if err != nil {
			return fmt.Errorf("解析内嵌APK失败: %v", err)
		}
		
		// 扫描内嵌APK
		ScanAPKData(embeddedReader)
		return nil
	}()
	
	return err
}

// 将APK读取到内存，传入
func ScanAPKData(apkReader *zip.Reader) error {
	//verifyApk(filePath) //20241230 临时取消签名验证 减小程序体积

	// 检测加固特征
	PackByLibSo(apkReader)

	// 根据参数控制安全检测
	if *ArgCheckRoot || *ArgCheckEmu || *ArgCheckDebug || *ArgCheckProxy {
		ScanAPKAnti(apkReader)
	}
	
	// 执行SDK检测
	if *ArgCheckSDK {
		SDKByLibSo(apkReader)
	}

	// 根据参数控制硬编码检测
	if *ArgCheckHardcode {
		ScanAPKHardCoded(apkReader)
	}

	// 根据参数控制证书检测
	if *ArgCheckCert {
		ScanAPKCertificate(apkReader)
	}
	
	return nil
}