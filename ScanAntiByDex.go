package main

import (
	"bytes"
	"fmt"
	"github.com/avast/apkparser"
	"os"
	"path"
)

// root检测常见 路径、字符串
var rootstringsCommonpaths = []string{
	"/cache/.disable_magisk",
	"/cache/magisk.log",
	"/cache/su",
	"/data/adb/ksu",
	"/data/adb/ksud",
	"/data/adb/magisk",
	"/data/adb/magisk.db",
	"/data/adb/magisk.img",
	"/data/adb/magisk_simple",
	"/data/local/bin/su",
	"/data/local/su",
	"/data/local/xbin/su",
	"/data/su",
	"/dev/.magisk.unblock",
	"/dev/com.koushikdutta.superuser.daemon/",
	"/dev/su",
	"/init.magisk.rc",
	"/sbin/.magisk",
	"/sbin/su",
	"/su/bin/su",
	"/system/app/Kinguser.apk",
	"/system/app/Superuser.apk",
	"/system/bin/.ext/su",
	"/system/bin/failsafe/su",
	"/system/bin/su",
	"/system/etc/init.d/99SuperSUDaemon",
	"/system/sbin/su",
	"/system/sd/xbin/su",
	"/system/usr/we-need-root/su",
	"/system/xbin/busybox",
	"/system/xbin/daemonsu",
	"/system/xbin/ku.sud",
	"/system/xbin/su",
	"/vendor/bin/su",
	"Kinguser.apk",  //某些检测会将路径字符串分开
	"Superuser.apk", //某些检测会将路径字符串分开
	"/system/xbin/", //某些检测会将路径字符串分开
	"/vendor/bin/",  //某些检测会将路径字符串分开
}

// root检测常见 apk 包名字符串
var rootstringsmanagementApp = []string{
	"com.chelpus.lackypatch",
	"com.dimonvideo.luckypatcher",
	"com.koushikdutta.rommanager",
	"com.koushikdutta.rommanager.license",
	"com.koushikdutta.superuser",
	"com.noshufou.android.su",
	"com.noshufou.android.su.elite",
	"com.ramdroid.appquarantine",
	"com.ramdroid.appquarantinepro",
	"com.thirdparty.superuser",
	"com.topjohnwu.magisk",
	"com.yellowes.su",
	"eu.chainfire.supersu",
	"me.weishu.kernelsu",
}

// 模拟器检查
var emulatorStrings = []string{
	"tel:123456",
	"test-keys",
	"goldfish",
	"android-test",
	"000000000000000",
	"/dev/socket/qemud",
	"/dev/qemu_pipe",
	"/dev/qemu_trace",
}

// 反调试检测
var DebugStrings = []string{
	"checkFridaRunningProcesses",          //getSystemService("activity").getRunningServices(300)
	"checkRunningProcesses",               //"frida-server"、"fridaserver"
	"checkRunningServices",                //supersu、superuser进程名检测
	"threadCpuTimeNanos",                  // cpu计算时间差检测是否被调试
	"TamperingWithJavaRuntime",            //篡改Java运行时
	"com.android.internal.os.ZygoteInit",  //篡改Java运行时
	"com.saurik.substrate.MS$2",           //篡改Java运行时
	"de.robv.android.xposed.XposedBridge", //篡改Java运行时
	"detectBypassSSL",
}

func ScanDexAnti(dexData []byte, filePath string) {
	// 搜索dex文件中是否包含root检测特征字符串
	for _, str := range rootstringsCommonpaths {
		if bytes.Contains(dexData, []byte(str)) {
			fmt.Printf("发现ROOT检测特征   [dex]: %s->%s\n", str, filePath)
			break
		}
	}
	// 搜索dex文件中是否包含模拟器检测特征字符串
	for _, str := range emulatorStrings {
		if bytes.Contains(dexData, []byte(str)) {
			fmt.Printf("发现模拟器检测特征 [dex]: %s->%s\n", str, filePath)
			break
		}
	}
	// 搜索dex文件中是否包含运行篡改检测特征函数
	for _, str := range DebugStrings {
		if bytes.Contains(dexData, []byte(str)) {
			fmt.Printf("发现反调试检测特征 [dex]: %s->%s\n", str, filePath)
			//break //因为包含了反调试、反sslbypass、反java运行篡改 不做直接跳出
		}
	}
}

func ScanAPKAnti(apkpath string) bool {
	//解析apk文件
	apkReader, err := apkparser.OpenZip(apkpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	defer apkReader.Close()

	// 读取dex文件扫描
	for _, file := range apkReader.File {
		if path.Ext(file.Name) == ".dex" {
			//fmt.Printf("Scan %s\n", file.Name)
			var dexData = []byte{}
			dexData, err = file.ReadAll(1024 * 1024 * 100) //单个文件读取最大设置为100MB
			ScanDexAnti(dexData, file.Name)
		}
	}

	return true
}
