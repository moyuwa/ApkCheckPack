package certificate

import (
	"archive/zip"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"strings"
)

// Detect 检测APK中的证书信息
func Detect(apkReader *zip.Reader, isEmbedded bool) {
	mainIndent := "  - "
	secondaryIndent := "    - "
	tertiaryIndent := "      - "
	if isEmbedded {
		mainIndent = "    - "
		secondaryIndent = "      - "
		tertiaryIndent = "        - "
	}

	fmt.Printf("\n\n" + mainIndent + "证书扫描结果")

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
			fmt.Printf("\n"+secondaryIndent+"证书文件: %s", file.Name)

			rc, err := file.Open()
			if err != nil {
				fmt.Printf("\n"+tertiaryIndent+"无法打开证书文件: %v", err)
				continue
			}

			certData, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				fmt.Printf("\n"+tertiaryIndent+"读取证书文件失败: %v", err)
				continue
			}

			// 解析证书
			block, _ := pem.Decode(certData)
			if block == nil {
				// 如果不是PEM格式，尝试直接解析DER格式
				cert, err := x509.ParseCertificate(certData)
				if err != nil {
					fmt.Printf("\n"+tertiaryIndent+"解析证书失败: %v", err)
					continue
				}
				printCertInfo(cert, tertiaryIndent)
			} else {
				// PEM格式证书
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					fmt.Printf("\n"+tertiaryIndent+"解析证书失败: %v", err)
					continue
				}
				printCertInfo(cert, tertiaryIndent)
			}
		}
	}

	if !foundCert {
		fmt.Println("\n" + secondaryIndent + "未发现证书文件")
	}
}

func printCertInfo(cert *x509.Certificate, indent string) {
	detailIndent := "      - "
	if indent == "        - " {
		detailIndent = "          - "
	}

	fmt.Printf("\n"+detailIndent+"主题: %s\n", cert.Subject.String())
	fmt.Printf(detailIndent+"发行者: %s\n", cert.Issuer.String())
	fmt.Printf(detailIndent+"序列号: %X\n", cert.SerialNumber)
	fmt.Printf(detailIndent+"有效期: %s 至 %s\n",
		cert.NotBefore.Format("2006-01-02 15:04:05"),
		cert.NotAfter.Format("2006-01-02 15:04:05"))
	fmt.Printf(detailIndent+"签名算法: %s\n", cert.SignatureAlgorithm.String())

	if cert.NotAfter.Before(cert.NotBefore) {
		fmt.Println(detailIndent + "[!] 警告: 证书有效期异常")
	}

	fmt.Printf(detailIndent+"SHA1指纹: %X\n", cert.SubjectKeyId)

	keyUsageStr := ""
	if (cert.KeyUsage & x509.KeyUsageDigitalSignature) != 0 {
		keyUsageStr += " 数字签名"
	}
	if (cert.KeyUsage & x509.KeyUsageKeyEncipherment) != 0 {
		keyUsageStr += " 密钥加密"
	}
	if (cert.KeyUsage & x509.KeyUsageCertSign) != 0 {
		keyUsageStr += " 证书签名"
	}
	if keyUsageStr != "" {
		fmt.Printf(detailIndent+"密钥用途:%s\n", keyUsageStr)
	}
}
