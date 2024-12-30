package main

//20241230 临时取消签名验证 减小程序体积
//
//import (
//	"fmt"
//	"github.com/avast/apkverifier"
//	"os"
//)
//
//func verifyApk(apkpath string) bool {
//	//读取配置
//	res, err := apkverifier.Verify(apkpath, nil)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Verification failed: %s\n", err.Error())
//		return false
//	}
//	//判断是否为V1版本
//	if res.SigningSchemeId == 1 {
//		fmt.Printf("Verification scheme used: v%d 版本签名,有Janus漏洞!\n", res.SigningSchemeId)
//	} else {
//		fmt.Printf("Verification scheme used: v%d 版本签名,无Janus漏洞!\n", res.SigningSchemeId)
//	}
//	//输出相信信息
//	cert, _ := apkverifier.PickBestApkCert(res.SignerCerts)
//	if cert == nil {
//		fmt.Printf("No certificate found.\n")
//	} else {
//		fmt.Println(cert)
//	}
//
//	return true
//}
