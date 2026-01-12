package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"
)

// 通用检测特征结构
type DetectionPattern struct {
    Pattern     string
    Description string
}

// Root文件路径检测特征
var RootFilePatterns = []DetectionPattern{
	{Pattern: "/cache/.disable_magisk", Description: "Magisk禁用标志文件"},
	{Pattern: "/cache/magisk.log", Description: "Magisk日志文件"},
	{Pattern: "/cache/su", Description: "SuperSU残留文件"},
	{Pattern: "/data/adb/ksu", Description: "KernelSU安装目录"},
	{Pattern: "/data/adb/ksud", Description: "KernelSU守护进程"},
	{Pattern: "/data/adb/magisk", Description: "Magisk主目录"},
	{Pattern: "/data/adb/magisk.db", Description: "Magisk策略数据库"},
	{Pattern: "/data/adb/magisk.img", Description: "Magisk镜像文件"},
	{Pattern: "/data/adb/magisk_simple", Description: "Magisk简化模式标记"},
	{Pattern: "/data/local/bin/su", Description: "SU二进制文件常见路径"},
	{Pattern: "/data/local/su", Description: "SU二进制备用路径"},
	{Pattern: "/data/local/xbin/su", Description: "Xposed框架SU路径"},
	{Pattern: "/data/su", Description: "SU配置目录"},
	{Pattern: "/dev/.magisk.unblock", Description: "Magisk解锁标记"},
	{Pattern: "/dev/com.koushikdutta.superuser.daemon/", Description: "Superuser守护进程socket"},
	{Pattern: "/dev/su", Description: "SU设备节点"},
	{Pattern: "/init.magisk.rc", Description: "Magisk启动脚本"},
	{Pattern: "/sbin/.magisk", Description: "Magisk临时目录"},
	{Pattern: "/sbin/su", Description: "系统分区SU文件"},
	{Pattern: "/su/bin/su", Description: "Systemless SU路径"},
	{Pattern: "/system/app/Kinguser.apk", Description: "Kingroot安装包"},
	{Pattern: "/system/app/Superuser.apk", Description: "Superuser安装包"},
	{Pattern: "/system/bin/.ext/su", Description: "隐藏的SU文件"},
	{Pattern: "/system/bin/failsafe/su", Description: "故障安全模式SU"},
	{Pattern: "/system/bin/su", Description: "系统内置SU"},
	{Pattern: "/system/etc/init.d/99SuperSUDaemon", Description: "SuperSU守护进程脚本"},
	{Pattern: "/system/sbin/su", Description: "系统分区备用SU路径"},
	{Pattern: "/system/sd/xbin/su", Description: "SD卡扩展SU路径"},
	{Pattern: "/system/usr/we-need-root/su", Description: "特殊目录SU文件"},
	{Pattern: "/system/xbin/busybox", Description: "BusyBox工具（常与root关联）"},
	{Pattern: "/system/xbin/daemonsu", Description: "SuperSU守护进程"},
	{Pattern: "/system/xbin/ku.sud", Description: "Kinguser守护进程"},
	{Pattern: "/system/xbin/su", Description: "常见SU存放路径"},
	{Pattern: "/vendor/bin/su", Description: "厂商分区SU文件"},
	{Pattern: "Kinguser.apk", Description: "Kingroot安装包（分段检测）"},
	{Pattern: "Superuser.apk", Description: "Superuser安装包（分段检测）"},
	{Pattern: "/system/xbin/", Description: "常见Root工具目录（分段检测）"},
	{Pattern: "/vendor/bin/", Description: "厂商Root工具目录（分段检测）"},
}

// Root管理应用包名检测特征
var RootAppPatterns = []DetectionPattern{
    {Pattern: "com.chelpus.lackypatch", Description: "Lucky Patcher工具"},
    {Pattern: "com.dimonvideo.luckypatcher", Description: "Lucky Patcher官方包名"},
    {Pattern: "com.koushikdutta.rommanager", Description: "Rom Manager刷机工具"},
    {Pattern: "com.koushikdutta.rommanager.license", Description: "Rom Manager许可证"},
    {Pattern: "com.koushikdutta.superuser", Description: "Koush's Superuser"},
    {Pattern: "com.noshufou.android.su", Description: "Superuser官方包名"},
    {Pattern: "com.noshufou.android.su.elite", Description: "Superuser付费版"},
    {Pattern: "com.ramdroid.appquarantine", Description: "应用隔离工具（需root）"},
    {Pattern: "com.ramdroid.appquarantinepro", Description: "专业版应用隔离"},
    {Pattern: "com.thirdparty.superuser", Description: "第三方Superuser应用"},
    {Pattern: "com.topjohnwu.magisk", Description: "Magisk Manager"},
    {Pattern: "com.yellowes.su", Description: "早期SuperSU变种"},
    {Pattern: "eu.chainfire.supersu", Description: "Chainfire SuperSU"},
    {Pattern: "me.weishu.kernelsu", Description: "KernelSU管理器"},
    {Pattern: "com.kingroot.kinguser", Description: "KingRoot主程序"},
    {Pattern: "com.kingoapp.root", Description: "KingoRoot"},
    {Pattern: "me.phh.superuser", Description: "PHH's Superuser"},
    {Pattern: "com.apusapps.browser.module.root", Description: "浏览器类Root工具"},
    {Pattern: "io.github.vvb2060.magisk", Description: "Magisk衍生版本"},
    {Pattern: "com.topjohnwu.magisk.pro", Description: "Magisk专业版"},
    {Pattern: "de.robv.android.xposed.installer", Description: "Xposed安装器"},
    {Pattern: "org.meowcat.edxposed.manager", Description: "EdXposed管理器"},
    {Pattern: "me.weishu.exp", Description: "太极框架"},
    {Pattern: "com.speedsoftware.rootexplorer", Description: "Root Explorer"},
    {Pattern: "com.keramidas.TitaniumBackup", Description: "钛备份"},
    {Pattern: "com.joeykrim.rootcheck", Description: "Root验证工具"},
    {Pattern: "com.device.report", Description: "Root报告工具"},
    {Pattern: "com.qihoo.root", Description: "360 Root"},
    {Pattern: "com.dianxinos.optimizer.duplay", Description: "点心Root"},
    {Pattern: "com.geohot.towelroot", Description: "Towelroot漏洞利用"},
    {Pattern: "com.zachspong.temprootremove", Description: "临时Root工具"},
    {Pattern: "com.riru.core", Description: "Riru核心模块"},
    {Pattern: "com.github.topjohnwu.magisk.installer", Description: "Magisk安装器"},
    {Pattern: "com.alephzain.framaroot", Description: "Framaroot工具"},
    {Pattern: "org.chainfire.internet", Description: "Chainfire网络工具"},
}

// 模拟器检测特征
var EmulatorPatterns = []DetectionPattern{
    {Pattern: "tel:123456", Description: "模拟器默认电话号码"},
    {Pattern: "test-keys", Description: "测试版系统特征"},
    {Pattern: "goldfish", Description: "Android模拟器内核标识"},
    {Pattern: "android-test", Description: "测试环境标识"},
    {Pattern: "000000000000000", Description: "模拟器默认IMEI"},
    {Pattern: "/dev/socket/qemud", Description: "QEMU守护进程socket"},
    {Pattern: "/dev/qemu_pipe", Description: "QEMU管道通信接口"},
    {Pattern: "/dev/qemu_trace", Description: "QEMU调试跟踪接口"},
    {Pattern: "ro.kernel.qemu", Description: "QEMU内核属性标识"},
    {Pattern: "generic_x86", Description: "模拟器常见ABI"},
    {Pattern: "emulator", Description: "模拟器标识"},
    {Pattern: "ro.boot.virtual", Description: "虚拟化启动标识"},
    {Pattern: "ro.cloudbuild.software", Description: "云构建标识"},
    {Pattern: "ro.secureboot.lockstate", Description: "安全启动锁定状态异常"},
    {Pattern: "ro.cpu.virtual", Description: "虚拟化CPU标识"},
    {Pattern: "Build.PRODUCT=sdk_google", Description: "SDK构建标识"},
    {Pattern: "Build.MODEL=Android SDK built", Description: "SDK构建模型"},
    {Pattern: "Build.HARDWARE=goldfish", Description: "模拟器硬件标识"},
    {Pattern: "Build.FINGERPRINT=generic", Description: "通用指纹特征"},
    {Pattern: "Sensor.TYPE_SIGNIFICANT_MOTION", Description: "真机特有传感器检测"},
    {Pattern: "Sensor.TYPE_STEP_COUNTER", Description: "步进传感器检测"},
    {Pattern: "Sensor.TYPE_HEART_RATE", Description: "心率传感器检测"},
    {Pattern: "10.0.2.15", Description: "默认NAT网关IP"},
    {Pattern: "eth0", Description: "模拟器网络接口"},
    {Pattern: "dns.google", Description: "模拟器默认DNS"},
    {Pattern: "debug.stagefright.ccode", Description: "Stagefright框架特征"},
    {Pattern: "ro.kernel.android.checkjni", Description: "JNI检查模式"},
    {Pattern: "ro.boot.selinux=disabled", Description: "SELinux禁用状态"},
    {Pattern: "hasQemuSocket", Description: "QEMU socket检测函数"},
    {Pattern: "hasQemuPipe", Description: "QEMU管道检测函数"},
    {Pattern: "getEmulatorQEMUKernel", Description: "QEMU内核属性检测函数"},
    {Pattern: "Landroid/os/SystemProperties;->get(Ljava/lang/String;)", Description: "系统属性调用模式"},
}

// 反调试检测特征
var DebugPatterns = []DetectionPattern{
    {Pattern: "checkFridaRunningProcesses", Description: "Frida进程检测"},
    {Pattern: "checkRunningProcesses", Description: "检测frida-server进程"},
    {Pattern: "checkRunningServices", Description: "检测supersu/superuser服务"},
    {Pattern: "treadCpuTimeNanos", Description: "CPU时间差调试检测"},
    {Pattern: "TamperingWithJavaRuntime", Description: "Java运行时篡改检测"},
    {Pattern: "com.android.internal.os.ZygoteInit", Description: "Zygote初始化检测"},
    {Pattern: "com.saurik.substrate.MS$2", Description: "Substrate框架检测"},
    {Pattern: "de.robv.android.xposed.XposedBridge", Description: "Xposed框架检测"},
    {Pattern: "detectBypassSSL", Description: "SSL证书绕过检测"},
    {Pattern: "Landroid/os/Debug;->isDebuggerConnected()Z", Description: "调试器连接检测"},
    {Pattern: ":27042", Description: "Frida默认端口检测"},
    {Pattern: ":23946", Description: "ADB调试端口检测"},
    {Pattern: "frida-gadget", Description: "Frida工具检测"},
    {Pattern: "libfrida.so", Description: "Frida库文件检测"},
    {Pattern: "XposedBridge.jar", Description: "Xposed桥接文件检测"},
    {Pattern: "EdXposed", Description: "EdXposed框架检测"},
    {Pattern: "frida-server", Description: "Frida服务器进程检测"},
    {Pattern: "android_server", Description: "IDA调试服务器检测"},
    {Pattern: "gdbserver", Description: "GDB调试服务器检测"},
    {Pattern: "ro.debuggable", Description: "系统调试属性检测"},
    {Pattern: "service.adb.root", Description: "ADB Root服务检测"},
    {Pattern: "XposedInstaller", Description: "Xposed安装器检测"},
    {Pattern: "Magisk", Description: "Magisk框架检测"},
    {Pattern: "LSPosed", Description: "LSPosed框架检测"},
    {Pattern: "ptrace", Description: "Ptrace调试检测"},
    {Pattern: "/proc/self/status", Description: "TracerPid状态检测"},
    {Pattern: "libsubstrate.so", Description: "Substrate库文件检测"},
    {Pattern: "com.saurik.substrate", Description: "Substrate框架检测"},
    {Pattern: "sslunpinning", Description: "SSL Unpinning检测"},
    {Pattern: "JustTrustMe", Description: "SSL证书绕过模块检测"},
    {Pattern: "/data/data/de.robv.android.xposed.installer/conf/modules.list", Description: "Xposed模块列表检测"},
}

// 代理检测特征列表
var ProxyPatterns = []DetectionPattern{
    {Pattern: "Lokhttp3/Proxy;->NO_PROXY:Lokhttp3/Proxy;", Description: "OkHttp显式禁用代理"},
    {Pattern: "Lokhttp3/OkHttpClient$Builder;->proxy(Lokhttp3/Proxy;)Lokhttp3/OkHttpClient$Builder;", Description: "OkHttp代理设置方法调用"},
    {Pattern: "Lokhttp3/internal/proxy/NullProxySelector;", Description: "OkHttp空代理选择器"},
    {Pattern: "Ljava/lang/System;->getProperty(Ljava/lang/String;)Ljava/lang/String;", Description: "系统代理属性访问"},
    {Pattern: "Landroid/net/Proxy;->getDefaultProxy()Landroid/net/Proxy;", Description: "获取默认代理"},
    {Pattern: "Landroid/net/Proxy;->getHost()Ljava/lang/String;", Description: "获取代理主机"},
    {Pattern: "Landroid/net/Proxy;->getPort()I", Description: "获取代理端口"},
    {Pattern: "Landroid/net/ConnectivityManager;->getActiveNetworkInfo", Description: "获取活动网络信息"},
    {Pattern: "Landroid/net/NetworkInfo;->getType", Description: "获取网络类型"},
    {Pattern: "Ljavax/net/ssl/X509TrustManager;", Description: "自定义证书信任管理器"},
    {Pattern: "Ljavax/net/ssl/SSLContext;->init", Description: "初始化SSL上下文"},
    {Pattern: "VPNService", Description: "VPN服务检测"},
    {Pattern: "NetworkCapabilities.TRANSPORT_VPN", Description: "VPN网络能力检测"},
    {Pattern: "isVpnUsed", Description: "VPN使用检测"},
}

func ScanDexAnti(dexData []byte, filePath string) {
	// 根据参数控制ROOT检测
	if *ArgCheckRoot {
		// 检查Root文件路径
		for _, pattern := range RootFilePatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", 
					filePath, pattern.Pattern, pattern.Description)
				allRootResults = append(allRootResults, result)
			}
		}
		// 检查Root应用包名
		for _, pattern := range RootAppPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", 
					filePath, pattern.Pattern, pattern.Description)
				allRootResults = append(allRootResults, result)
			}
		}
	}

	// 根据参数控制模拟器检测
	if *ArgCheckEmu {
		for _, pattern := range EmulatorPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", 
					filePath, pattern.Pattern, pattern.Description)
				allEmuResults = append(allEmuResults, result)
			}
		}
	}

	// 根据参数控制反调试检测
	if *ArgCheckDebug {
		for _, pattern := range DebugPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", 
					filePath, pattern.Pattern, pattern.Description)
				allDebugResults = append(allDebugResults, result)
			}
		}
	}

	// 根据参数控制代理检测
	if *ArgCheckProxy {
		for _, pattern := range ProxyPatterns {
			if bytes.Contains(dexData, []byte(pattern.Pattern)) {
				result := fmt.Sprintf("    %s -> %s (%s)", 
					filePath, pattern.Pattern, pattern.Description)
				allProxyResults = append(allProxyResults, result)
			}
		}
	}
}

// 全局变量用于收集所有DEX文件的检测结果
var (
	allRootResults  []string
	allEmuResults   []string
	allDebugResults []string
	allProxyResults []string
)

// 清除所有检测结果
func clearAntiResults() {
	// 预分配足够大的切片容量，避免频繁的内存重新分配
	allRootResults = make([]string, 0, 500)
	allEmuResults = make([]string, 0, 500)
	allDebugResults = make([]string, 0, 500)
	allProxyResults = make([]string, 0, 500)
}

// 计算最长特征长度，用于滑动窗口处理
func getMaxPatternLength() int {
	maxLen := 0
	
	// 检查所有特征的长度
	for _, pattern := range RootFilePatterns {
		if len(pattern.Pattern) > maxLen {
			maxLen = len(pattern.Pattern)
		}
	}
	
	for _, pattern := range RootAppPatterns {
		if len(pattern.Pattern) > maxLen {
			maxLen = len(pattern.Pattern)
		}
	}
	
	for _, pattern := range EmulatorPatterns {
		if len(pattern.Pattern) > maxLen {
			maxLen = len(pattern.Pattern)
		}
	}
	
	for _, pattern := range DebugPatterns {
		if len(pattern.Pattern) > maxLen {
			maxLen = len(pattern.Pattern)
		}
	}
	
	for _, pattern := range ProxyPatterns {
		if len(pattern.Pattern) > maxLen {
			maxLen = len(pattern.Pattern)
		}
	}
	
	return maxLen
}

// 分块读取并检测文件内容，避免一次性加载大文件到内存
func ScanFileByChunk(fileReader io.Reader, filePath string, maxSize int64) {
	// 获取最长特征长度
	maxPatternLen := getMaxPatternLength()
	
	// 设置块大小为8MB，加上最长特征长度作为重叠部分
	chunkSize := int64(8 * 1024 * 1024)
	overlapSize := int64(maxPatternLen) // 转换为int64类型
	
	// 创建缓冲区
	buffer := make([]byte, chunkSize+overlapSize)
	prevBuffer := make([]byte, overlapSize)
	
	// 限制读取大小
	limitReader := io.LimitReader(fileReader, maxSize)
	
	for {
		// 读取数据
		n, err := limitReader.Read(buffer[overlapSize:])
		if n == 0 {
			break
		}
		
		// 将上一次的末尾部分复制到当前缓冲区的开头
		copy(buffer, prevBuffer)
		
		// 更新当前处理的数据长度
		dataLen := overlapSize + int64(n) // 转换n为int64类型
		
		// 检测当前块中的特征
		ScanDexAnti(buffer[:dataLen], filePath)
		
		// 保存当前块的末尾部分，用于下一次检测
		if dataLen > overlapSize {
			copy(prevBuffer, buffer[dataLen-overlapSize:dataLen])
		} else {
			copy(prevBuffer, buffer[:dataLen])
			// 不足重叠大小，填充剩余部分
			if dataLen < overlapSize {
				for i := dataLen; i < overlapSize; i++ {
					prevBuffer[i] = 0
				}
			}
		}
		
		if err != nil {
			break
		}
	}
}

func ScanAPKAnti(apkReader *zip.Reader) bool {
    // 清除之前的检测结果
    clearAntiResults()
    
    // 读取dex文件扫描
    for _, file := range apkReader.File {
        if path.Ext(file.Name) == ".dex" {
            fileReader, err := file.Open()
            if err != nil {
                fmt.Println(err)
                continue
            }
            
            // 使用main包中的maxsize参数
            maxSize := int64(*ArgMaxSize) * 1024 * 1024 // 转换为字节单位
            
            // 分块读取和处理文件，避免一次性加载大文件到内存
            ScanFileByChunk(fileReader, file.Name, maxSize)
            fileReader.Close()
        }
    }
    
    // 输出结果
    // 根据是否是内嵌APK调整缩进级别
    mainIndent := "  - "
    secondaryIndent := "    - "
    tertiaryIndent := "      - "
    if isEmbeddedAPK {
        mainIndent = "    - "
        secondaryIndent = "      - "
        tertiaryIndent = "        - "
    }
    
    fmt.Println("\n" + mainIndent + "安全检测特征扫描结果")
    
    // 统计所有安全检测结果的数量
    totalResults := 0
    if *ArgCheckRoot {
        totalResults += len(allRootResults)
    }
    if *ArgCheckEmu {
        totalResults += len(allEmuResults)
    }
    if *ArgCheckDebug {
        totalResults += len(allDebugResults)
    }
    if *ArgCheckProxy {
        totalResults += len(allProxyResults)
    }
    
    // 如果没有任何安全检测特征，输出提示
    if totalResults == 0 {
        fmt.Println(secondaryIndent + "未发现安全检测特征")
    } else {
        if *ArgCheckRoot && len(allRootResults) > 0 {
            fmt.Println(secondaryIndent + "ROOT检测特征")
            for _, result := range allRootResults {
                // 移除行首的空格，然后添加适当的缩进
                trimmedResult := strings.TrimSpace(result)
                fmt.Printf(tertiaryIndent + "%s\n", trimmedResult)
            }
        }

        if *ArgCheckEmu && len(allEmuResults) > 0 {
            fmt.Println(secondaryIndent + "模拟器检测特征")
            for _, result := range allEmuResults {
                // 移除行首的空格，然后添加适当的缩进
                trimmedResult := strings.TrimSpace(result)
                fmt.Printf(tertiaryIndent + "%s\n", trimmedResult)
            }
        }

        if *ArgCheckDebug && len(allDebugResults) > 0 {
            fmt.Println(secondaryIndent + "反调试检测特征")
            for _, result := range allDebugResults {
                // 移除行首的空格，然后添加适当的缩进
                trimmedResult := strings.TrimSpace(result)
                fmt.Printf(tertiaryIndent + "%s\n", trimmedResult)
            }
        }

        if *ArgCheckProxy && len(allProxyResults) > 0 {
            fmt.Println(secondaryIndent + "代理检测特征")
            for _, result := range allProxyResults {
                // 移除行首的空格，然后添加适当的缩进
                trimmedResult := strings.TrimSpace(result)
                fmt.Printf(tertiaryIndent + "%s\n", trimmedResult)
            }
        }
    }
    
    // fmt.Println("\n================================================================")

    return true
}