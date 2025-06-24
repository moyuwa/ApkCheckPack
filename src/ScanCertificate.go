package main

import (
	"archive/zip"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"strings"
)

// 扫描APK证书信息
func ScanAPKCertificate(apkReader *zip.Reader) {
	fmt.Printf("\n===================== 证书扫描结果 =====================\n")
	
	// 定义证书文件后缀
	certExtensions := []string{
		".CRT",
		".CER",
		".PEM",
		".DER",
		".P12",
		".PFX",
		".RSA",
		".DSA",
	}

	foundCert := false
	for _, file := range apkReader.File {
		// 检查文件是否是证书
		isMatch := false
		upperName := strings.ToUpper(file.Name)
		for _, ext := range certExtensions {
			if strings.HasSuffix(upperName, ext) {
				isMatch = true
				break
			}
		}

		if isMatch {
			foundCert = true
			fmt.Printf("\n[证书文件: %s]\n", file.Name)
			
			rc, err := file.Open()
			if err != nil {
				fmt.Printf("    无法打开证书文件: %v\n", err)
				continue
			}
			defer rc.Close()
			
			certData, err := io.ReadAll(rc)
			if err != nil {
				fmt.Printf("    读取证书文件失败: %v\n", err)
				continue
			}
			
			// 解析证书
			block, _ := pem.Decode(certData)
			if block == nil {
				// 如果不是PEM格式，尝试直接解析DER格式
				cert, err := x509.ParseCertificate(certData)
				if err != nil {
					fmt.Printf("    解析证书失败: %v\n", err)
					continue
				}
				
				printCertificateInfo(cert)
			} else {
				// PEM格式证书
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					fmt.Printf("    解析证书失败: %v\n", err)
					continue
				}
				
				printCertificateInfo(cert)
			}
		}
	}
	
	if !foundCert {
		fmt.Println("\n[!] 未发现证书文件")
	}
	
	// fmt.Printf("\n================================================================\n")
}

// 打印证书详细信息
func printCertificateInfo(cert *x509.Certificate) {
	fmt.Printf("    主题: %s\n", cert.Subject.String())
	fmt.Printf("    发行者: %s\n", cert.Issuer.String())
	fmt.Printf("    序列号: %X\n", cert.SerialNumber)
	fmt.Printf("    有效期: %v 至 %v\n", cert.NotBefore.Format("2006-01-02 15:04:05"), cert.NotAfter.Format("2006-01-02 15:04:05"))
	fmt.Printf("    签名算法: %s\n", cert.SignatureAlgorithm.String())
	
	// 检查证书是否过期
	if cert.NotAfter.Before(cert.NotBefore) {
		fmt.Println("    [!] 警告: 证书有效期异常")
	}
	
	// 打印证书指纹
	fmt.Printf("    SHA1指纹: %X\n", cert.SubjectKeyId)
	
	// 打印密钥用途
	fmt.Printf("    密钥用途:")
	if (cert.KeyUsage & x509.KeyUsageDigitalSignature) != 0 {
		fmt.Printf(" 数字签名")
	}
	if (cert.KeyUsage & x509.KeyUsageKeyEncipherment) != 0 {
		fmt.Printf(" 密钥加密")
	}
	if (cert.KeyUsage & x509.KeyUsageCertSign) != 0 {
		fmt.Printf(" 证书签名")
	}
	fmt.Println()
}