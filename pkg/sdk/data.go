package sdk

// SDKLangInfo 语言相关的SDK信息
type SDKLangInfo struct {
	Label   string
	DevTeam string
}

// SDKRule 第三方SDK规则
type SDKRule struct {
	Soname string
	Zh     SDKLangInfo
	En     SDKLangInfo
}

// GetLabel 获取标签，优先中文
func (s *SDKRule) GetLabel() string {
	if s.Zh.Label != "" {
		return s.Zh.Label
	}
	return s.En.Label
}

// GetTeam 获取团队名称，优先中文
func (s *SDKRule) GetTeam() string {
	if s.Zh.DevTeam != "" {
		return s.Zh.DevTeam
	}
	return s.En.DevTeam
}

// SDKRules 内置的SDK规则列表
var SDKRules = []SDKRule{
	{
		Soname: `libopencv_java3.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `librxffmpeg-player.so`,
		Zh: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
		En: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
	},
	{
		Soname: `libmpbase.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libxray.so`,
		Zh: SDKLangInfo{
			Label:   `Xray Core`,
			DevTeam: `Project X Community`,
		},
		En: SDKLangInfo{
			Label:   `Xray Core`,
			DevTeam: `Project X Community`,
		},
	},
	{
		Soname: `libxamarin-app.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libsqlite3requery.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libmlkitcommonpipeline.so`,
		Zh: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `liblspd.so`,
		Zh: SDKLangInfo{
			Label:   `LSPatch`,
			DevTeam: `LSPosed`,
		},
		En: SDKLangInfo{
			Label:   `LSPatch`,
			DevTeam: `LSPosed`,
		},
	},
	{
		Soname: `libndkbitmap.so`,
		Zh: SDKLangInfo{
			Label:   `烈焰弹幕使`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `Flame Barrage Master`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libSipCryptor.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
		En: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
	},
	{
		Soname: `libweex_xr_plugin.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libP2PController.so`,
		Zh: SDKLangInfo{
			Label:   `MobileIMSDK`,
			DevTeam: `JackJiang`,
		},
		En: SDKLangInfo{
			Label:   `MobileIMSDK`,
			DevTeam: `JackJiang`,
		},
	},
	{
		Soname: `liblynxdevtool.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libchrome_android_linker.so`,
		Zh: SDKLangInfo{
			Label:   `Android Crazy Linker`,
			DevTeam: `The Chromium Authors`,
		},
		En: SDKLangInfo{
			Label:   `Android Crazy Linker`,
			DevTeam: `The Chromium Authors`,
		},
	},
	{
		Soname: `libnative-tianmu-common.so`,
		Zh: SDKLangInfo{
			Label:   `Tianmu SDK`,
			DevTeam: `杭州艾狄墨搏信息服务有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Tianmu SDK`,
			DevTeam: `杭州艾狄墨搏信息服务有限公司`,
		},
	},
	{
		Soname: `libduiutils.so`,
		Zh: SDKLangInfo{
			Label:   `思必驰`,
			DevTeam: `AISPEECH`,
		},
		En: SDKLangInfo{
			Label:   `Sibichi`,
			DevTeam: `AISPEECH`,
		},
	},
	{
		Soname: `libglog.so`,
		Zh: SDKLangInfo{
			Label:   `glog`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `glog`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libpine.so`,
		Zh: SDKLangInfo{
			Label:   `Pine`,
			DevTeam: `canyie`,
		},
		En: SDKLangInfo{
			Label:   `Pine`,
			DevTeam: `canyie`,
		},
	},
	{
		Soname: `libDexHelper-x86.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libffmpeg_mediametadataretriever_jni.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libpangleflipped.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libvrtoolkit.so`,
		Zh: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libwechatns.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libcfcaMLog.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
		En: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
	},
	{
		Soname: `libsecenh_x86.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA Reinforcement`,
			DevTeam: `CFCA`,
		},
		En: SDKLangInfo{
			Label:   `CFCA 加固`,
			DevTeam: `CFCA`,
		},
	},
	{
		Soname: `libovpn3.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libmonochrome_64.so`,
		Zh: SDKLangInfo{
			Label:   `libmonochrome`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `libmonochrome`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `libbd_etts.so`,
		Zh: SDKLangInfo{
			Label:   `离线语音合成 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Offline Speech Synthesis SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libovpnexec.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libtaro_native_bridge.so`,
		Zh: SDKLangInfo{
			Label:   `Taro`,
			DevTeam: `京东`,
		},
		En: SDKLangInfo{
			Label:   `Taro`,
			DevTeam: `京东`,
		},
	},
	{
		Soname: `libappmodules.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libwebviewchromium.huawei.so`,
		Zh: SDKLangInfo{
			Label:   `HUAWEI WebView`,
			DevTeam: `Chromium, HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HUAWEI WebView`,
			DevTeam: `Chromium, HUAWEI`,
		},
	},
	{
		Soname: `libtoyger.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libfdog.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libGCloudVoice.so`,
		Zh: SDKLangInfo{
			Label:   `游戏语音 GVoice`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Game Voice GVoice`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmodft2.so`,
		Zh: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libyoga.so`,
		Zh: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libunreal.so`,
		Zh: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
		En: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
	},
	{
		Soname: `libOvenPlayerCore.so`,
		Zh: SDKLangInfo{
			Label:   `OvenPlayer`,
			DevTeam: `AirenSoft`,
		},
		En: SDKLangInfo{
			Label:   `OvenPlayer`,
			DevTeam: `AirenSoft`,
		},
	},
	{
		Soname: `libijkffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libdrc.so`,
		Zh: SDKLangInfo{
			Label:   `libdrc`,
			DevTeam: `Mema Hacking`,
		},
		En: SDKLangInfo{
			Label:   `libdrc`,
			DevTeam: `Mema Hacking`,
		},
	},
	{
		Soname: `libp7zip.so`,
		Zh: SDKLangInfo{
			Label:   `AndroidP7zip`,
			DevTeam: `hzy3774`,
		},
		En: SDKLangInfo{
			Label:   `AndroidP7zip`,
			DevTeam: `hzy3774`,
		},
	},
	{
		Soname: `libsls_producer.so`,
		Zh: SDKLangInfo{
			Label:   `日志服务 SLS`,
			DevTeam: `Aliyun`,
		},
		En: SDKLangInfo{
			Label:   `Log Service SLS`,
			DevTeam: `Aliyun`,
		},
	},
	{
		Soname: `libhwcpipe.so`,
		Zh: SDKLangInfo{
			Label:   `HWCPipe`,
			DevTeam: `ARM-software`,
		},
		En: SDKLangInfo{
			Label:   `HWCPipe`,
			DevTeam: `ARM-software`,
		},
	},
	{
		Soname: `liblbtrtc.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libDnsChecker.so`,
		Zh: SDKLangInfo{
			Label:   `DnsChecker`,
			DevTeam: `olof `,
		},
		En: SDKLangInfo{
			Label:   `DnsChecker`,
			DevTeam: `olof `,
		},
	},
	{
		Soname: `libbd_facecollect_unifylicense.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libyuvutil.so`,
		Zh: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libsophix.so`,
		Zh: SDKLangInfo{
			Label:   `阿里移动热修复`,
			DevTeam: `Aliyun`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Mobile Hotfix`,
			DevTeam: `Aliyun`,
		},
	},
	{
		Soname: `libapplovin-native-crash-reporter.so`,
		Zh: SDKLangInfo{
			Label:   `AppLovin`,
			DevTeam: `AppLovin`,
		},
		En: SDKLangInfo{
			Label:   `AppLovin`,
			DevTeam: `AppLovin`,
		},
	},
	{
		Soname: `libreact_codegen_yrncore.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `librtmp-jni.so`,
		Zh: SDKLangInfo{
			Label:   `librtmp`,
			DevTeam: `ant-media`,
		},
		En: SDKLangInfo{
			Label:   `librtmp`,
			DevTeam: `ant-media`,
		},
	},
	{
		Soname: `libmediastreamer_voip.so`,
		Zh: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
		En: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
	},
	{
		Soname: `libauth.so`,
		Zh: SDKLangInfo{
			Label:   `libauth`,
			DevTeam: `Bitauth`,
		},
		En: SDKLangInfo{
			Label:   `libauth`,
			DevTeam: `Bitauth`,
		},
	},
	{
		Soname: `libapmart.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯客户端性能分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Client Performance Analysis`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libMNN_dl.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libglog_init.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libtdmqimei.so`,
		Zh: SDKLangInfo{
			Label:   `Qimei SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Qimei SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libYoutuOcrJniApi.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmain.so`,
		Zh: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `libagora_content_inspect_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libRDTAPIsT.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libvdog-x86.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libruntimeexecutor.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libandroid-tree-sitter.so`,
		Zh: SDKLangInfo{
			Label:   `android-tree-sitter`,
			DevTeam: `AndroidIDEOfficial`,
		},
		En: SDKLangInfo{
			Label:   `android-tree-sitter`,
			DevTeam: `AndroidIDEOfficial`,
		},
	},
	{
		Soname: `libAliNNPython.so`,
		Zh: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libstretch.so`,
		Zh: SDKLangInfo{
			Label:   `Stretch`,
			DevTeam: `vislyhq（visly）`,
		},
		En: SDKLangInfo{
			Label:   `Stretch`,
			DevTeam: `vislyhq（visly）`,
		},
	},
	{
		Soname: `libsnpe_adsp.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libyouga.so`,
		Zh: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbytertc_ffmpeg_audio_extension.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libMNNOpenCV.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libavif_android.so`,
		Zh: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
		En: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
	},
	{
		Soname: `libenet.so`,
		Zh: SDKLangInfo{
			Label:   `ENet`,
			DevTeam: `Lee Salzman`,
		},
		En: SDKLangInfo{
			Label:   `ENet`,
			DevTeam: `Lee Salzman`,
		},
	},
	{
		Soname: `libBugly-ext.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `librongcloud_xcrash.so`,
		Zh: SDKLangInfo{
			Label:   `融云 IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
		En: SDKLangInfo{
			Label:   `RongCloud IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
	},
	{
		Soname: `libwtbt828.so`,
		Zh: SDKLangInfo{
			Label:   `高德地图 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Amap SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libKF5ConfigQml_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libYUVDecoder.so`,
		Zh: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libhermes_executor.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libtxplayer.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云短视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Short Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmpaas_crypto.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libqalcodecwrapper.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云通信 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Communication SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libnesec.so`,
		Zh: SDKLangInfo{
			Label:   `网易易盾`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Yidun`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libMNN290.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libzucchini.so`,
		Zh: SDKLangInfo{
			Label:   `Zucchini`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Zucchini`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libTNN.so`,
		Zh: SDKLangInfo{
			Label:   `TNN`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `TNN`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libkoom-native.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libmatrix-pthreadhook.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libdownloadproxy.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云实时音视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Real-time Audio and Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtorrent4j.so`,
		Zh: SDKLangInfo{
			Label:   `Java 的 Torrent 库`,
			DevTeam: `Alden Torres`,
		},
		En: SDKLangInfo{
			Label:   `Torrent library for java`,
			DevTeam: `Alden Torres`,
		},
	},
	{
		Soname: `libtoyger_doc.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libhpplayvideo.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libFLAC.so`,
		Zh: SDKLangInfo{
			Label:   `Free Lossless Audio Codec`,
			DevTeam: `xiph`,
		},
		En: SDKLangInfo{
			Label:   `Free Lossless Audio Codec`,
			DevTeam: `xiph`,
		},
	},
	{
		Soname: `libtobEmbedPagEncrypt.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libspeechengine.so`,
		Zh: SDKLangInfo{
			Label:   `speechEngine`,
			DevTeam: `Audiokinetic Inc.`,
		},
		En: SDKLangInfo{
			Label:   `speechEngine`,
			DevTeam: `Audiokinetic Inc.`,
		},
	},
	{
		Soname: `libtxcopyrightedmedia.so`,
		Zh: SDKLangInfo{
			Label:   `正版曲库直通车`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Genuine Music Library Direct Train`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libquickjs.so`,
		Zh: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
		En: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
	},
	{
		Soname: `libCtaApiLib.so`,
		Zh: SDKLangInfo{
			Label:   `天翼账号认证`,
			DevTeam: `天翼数字生活科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `天翼账号认证`,
			DevTeam: `天翼数字生活科技有限公司`,
		},
	},
	{
		Soname: `libhta.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libopenvpn.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libv8jni.so`,
		Zh: SDKLangInfo{
			Label:   `V8`,
			DevTeam: `rubyjs`,
		},
		En: SDKLangInfo{
			Label:   `V8`,
			DevTeam: `rubyjs`,
		},
	},
	{
		Soname: `libvideoeffect.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libzeus_direct_dex.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libagora_audio_beauty_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libreactconfig.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbdface_sdk.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libanychatcore.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `libx264Encoder.so`,
		Zh: SDKLangInfo{
			Label:   `x264`,
			DevTeam: `VideoLAN`,
		},
		En: SDKLangInfo{
			Label:   `x264`,
			DevTeam: `VideoLAN`,
		},
	},
	{
		Soname: `liblynxtrace.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libdexpatch_x86.so`,
		Zh: SDKLangInfo{
			Label:   `DexPatch`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `DexPatch`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libJni_wgs2gcj.so`,
		Zh: SDKLangInfo{
			Label:   `wgs2gcj`,
			DevTeam: ``,
		},
		En: SDKLangInfo{
			Label:   `wgs2gcj`,
			DevTeam: ``,
		},
	},
	{
		Soname: `libpanglearmor.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libmatrix_mem_util.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbspatch-lib.so`,
		Zh: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains_v3.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libxcrash_dumper.so`,
		Zh: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
		En: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
	},
	{
		Soname: `libjsengine-platform.so`,
		Zh: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
		En: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
	},
	{
		Soname: `libjsenv.so`,
		Zh: SDKLangInfo{
			Label:   `jsenv`,
			DevTeam: `OpenAtom Foundation`,
		},
		En: SDKLangInfo{
			Label:   `jsenv`,
			DevTeam: `OpenAtom Foundation`,
		},
	},
	{
		Soname: `libarmavlink.so`,
		Zh: SDKLangInfo{
			Label:   `libARMavlink`,
			DevTeam: `Parrot`,
		},
		En: SDKLangInfo{
			Label:   `libARMavlink`,
			DevTeam: `Parrot`,
		},
	},
	{
		Soname: `libmp3lame.so`,
		Zh: SDKLangInfo{
			Label:   `LAME`,
			DevTeam: `The LAME Project`,
		},
		En: SDKLangInfo{
			Label:   `LAME`,
			DevTeam: `The LAME Project`,
		},
	},
	{
		Soname: `libpjsua2.so`,
		Zh: SDKLangInfo{
			Label:   `PJSIP`,
			DevTeam: `PJSIP`,
		},
		En: SDKLangInfo{
			Label:   `PJSIP`,
			DevTeam: `PJSIP`,
		},
	},
	{
		Soname: `libsoundtouch.so`,
		Zh: SDKLangInfo{
			Label:   `SoundTouch`,
			DevTeam: `Olli Parviainen`,
		},
		En: SDKLangInfo{
			Label:   `SoundTouch`,
			DevTeam: `Olli Parviainen`,
		},
	},
	{
		Soname: `libbytenn.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libbdSpilWakeup.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libbdaudioeffect.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libmnnface.so`,
		Zh: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libchromium_android_linker.so`,
		Zh: SDKLangInfo{
			Label:   `libchrome`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `libchrome`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `libliquidfun.so`,
		Zh: SDKLangInfo{
			Label:   `LiquidFun`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `LiquidFun`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libGVoicePlugin.so`,
		Zh: SDKLangInfo{
			Label:   `游戏语音 GVoice`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Game Voice GVoice`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `librlottie.so`,
		Zh: SDKLangInfo{
			Label:   `rlottie`,
			DevTeam: `Samsung`,
		},
		En: SDKLangInfo{
			Label:   `rlottie`,
			DevTeam: `Samsung`,
		},
	},
	{
		Soname: `libopencv_imgproc.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `libx264.so`,
		Zh: SDKLangInfo{
			Label:   `x264`,
			DevTeam: `VideoLAN`,
		},
		En: SDKLangInfo{
			Label:   `x264`,
			DevTeam: `VideoLAN`,
		},
	},
	{
		Soname: `libsgsecuritybodyso.version.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsnapdragon_services_adk.qti.so`,
		Zh: SDKLangInfo{
			Label:   `QESDK`,
			DevTeam: `Qualcomm`,
		},
		En: SDKLangInfo{
			Label:   `QESDK`,
			DevTeam: `Qualcomm`,
		},
	},
	{
		Soname: `libvdog.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libUniRapidJson.so`,
		Zh: SDKLangInfo{
			Label:   `UniRapidJson`,
			DevTeam: `takezoh`,
		},
		En: SDKLangInfo{
			Label:   `UniRapidJson`,
			DevTeam: `takezoh`,
		},
	},
	{
		Soname: `libzxingcpp_android.so`,
		Zh: SDKLangInfo{
			Label:   `zxing-cpp`,
			DevTeam: `zxing-cpp`,
		},
		En: SDKLangInfo{
			Label:   `zxing-cpp`,
			DevTeam: `zxing-cpp`,
		},
	},
	{
		Soname: `libagora-mpg123.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libbd_unifylicense.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libreact_codegen_FileAccessSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libwtcrypto.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云通信 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Communication SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmnnruntime.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `liblouis.so`,
		Zh: SDKLangInfo{
			Label:   `Liblouis`,
			DevTeam: `Liblouis`,
		},
		En: SDKLangInfo{
			Label:   `Liblouis`,
			DevTeam: `Liblouis`,
		},
	},
	{
		Soname: `libreactjscexecutor.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libdu.so`,
		Zh: SDKLangInfo{
			Label:   `数字联盟可信 ID SDK`,
			DevTeam: `北京数字联盟网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Digital Alliance Trusted ID SDK`,
			DevTeam: `北京数字联盟网络科技有限公司`,
		},
	},
	{
		Soname: `libstlport_shared.so`,
		Zh: SDKLangInfo{
			Label:   `SQLCipher`,
			DevTeam: `SQLCipher`,
		},
		En: SDKLangInfo{
			Label:   `SQLCipher`,
			DevTeam: `SQLCipher`,
		},
	},
	{
		Soname: `librealm_dart.so`,
		Zh: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
		En: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
	},
	{
		Soname: `libpaddle_full_api_shared.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libbarhopper_v2.so`,
		Zh: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libtaro_native_core.so`,
		Zh: SDKLangInfo{
			Label:   `Taro`,
			DevTeam: `京东`,
		},
		En: SDKLangInfo{
			Label:   `Taro`,
			DevTeam: `京东`,
		},
	},
	{
		Soname: `libdict-parser.so`,
		Zh: SDKLangInfo{
			Label:   `文字识别服务`,
			DevTeam: `有道`,
		},
		En: SDKLangInfo{
			Label:   `Text Recognition Service`,
			DevTeam: `有道`,
		},
	},
	{
		Soname: `liblynxdebugrouter.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `liblynx_helium.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libsecsdk.so`,
		Zh: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
		En: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
	},
	{
		Soname: `libnesec-x86.so`,
		Zh: SDKLangInfo{
			Label:   `网易易盾`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Yidun`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libksylive.so`,
		Zh: SDKLangInfo{
			Label:   `金山云短视频编辑 SDK`,
			DevTeam: `Kingsoft`,
		},
		En: SDKLangInfo{
			Label:   `Kingsoft Cloud Short Video Editing SDK`,
			DevTeam: `Kingsoft`,
		},
	},
	{
		Soname: `libdidi_secure.so`,
		Zh: SDKLangInfo{
			Label:   `滴滴 SDK`,
			DevTeam: `Didi`,
		},
		En: SDKLangInfo{
			Label:   `Didi SDK`,
			DevTeam: `Didi`,
		},
	},
	{
		Soname: `libmnn_wrapper.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `librxffmpeg-core.so`,
		Zh: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
		En: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
	},
	{
		Soname: `libtensorflowlite_flex_jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libjiagu.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libizuko.so`,
		Zh: SDKLangInfo{
			Label:   `Izuko`,
			DevTeam: `Absinthe`,
		},
		En: SDKLangInfo{
			Label:   `Izuko`,
			DevTeam: `Absinthe`,
		},
	},
	{
		Soname: `libmatrix-fd.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libDexHelper.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libagora_clear_vision_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libmpv.so`,
		Zh: SDKLangInfo{
			Label:   `LibMpv`,
			DevTeam: `Mpv`,
		},
		En: SDKLangInfo{
			Label:   `LibMpv`,
			DevTeam: `Mpv`,
		},
	},
	{
		Soname: `libagora-player-ffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libxcrash_dumper_e.so`,
		Zh: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
		En: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
	},
	{
		Soname: `libmilinkconnection.so`,
		Zh: SDKLangInfo{
			Label:   `小米游戏 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiaomi Game SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libhdog.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libtensorflowlite_jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libA3AEECD8.so`,
		Zh: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
		En: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
	},
	{
		Soname: `libreact_codegen_RNCSlider.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Slider`,
			DevTeam: `Callstack`,
		},
		En: SDKLangInfo{
			Label:   `React Native Slider`,
			DevTeam: `Callstack`,
		},
	},
	{
		Soname: `libyt_safe.so`,
		Zh: SDKLangInfo{
			Label:   `屹通 SDK`,
			DevTeam: `上海屹通信息科技发展有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Yitong SDK`,
			DevTeam: `上海屹通信息科技发展有限公司`,
		},
	},
	{
		Soname: `libEasyARPlayer.so`,
		Zh: SDKLangInfo{
			Label:   `EasyAR`,
			DevTeam: `视辰信息科技（上海）有限公司`,
		},
		En: SDKLangInfo{
			Label:   `EasyAR`,
			DevTeam: `视辰信息科技（上海）有限公司`,
		},
	},
	{
		Soname: `libncnn.so`,
		Zh: SDKLangInfo{
			Label:   `ncnn`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `ncnn`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbspatch_utils.so`,
		Zh: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libadmob-google.so`,
		Zh: SDKLangInfo{
			Label:   `Google AdMob`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google AdMob`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libvraudio_engine.so`,
		Zh: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libByteVC1_dec.so`,
		Zh: SDKLangInfo{
			Label:   `Broadcast SDK for Android`,
			DevTeam: `BytePlus`,
		},
		En: SDKLangInfo{
			Label:   `Broadcast SDK for Android`,
			DevTeam: `BytePlus`,
		},
	},
	{
		Soname: `libbangcle_risk.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libAkMatrixReverb.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libX86Bridge.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libexecoat.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libwtecdh.so`,
		Zh: SDKLangInfo{
			Label:   `米大师`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Mi Master`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbdEASRAndroid.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libaliresample.so`,
		Zh: SDKLangInfo{
			Label:   `阿里云短视频 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Cloud Short Video SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libbluestar_jni_android.so`,
		Zh: SDKLangInfo{
			Label:   `Bluestar`,
			DevTeam: `sandy98`,
		},
		En: SDKLangInfo{
			Label:   `Bluestar`,
			DevTeam: `sandy98`,
		},
	},
	{
		Soname: `libfolly_json.so`,
		Zh: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libmsaoaidauth.so`,
		Zh: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
		En: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
	},
	{
		Soname: `libcrypto.so`,
		Zh: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
		En: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
	},
	{
		Soname: `libliantiansec.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libvegamearm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `火山引擎云手机 SDK`,
			DevTeam: `Bytedance`,
		},
		En: SDKLangInfo{
			Label:   `Volcano Engine Cloud Phone SDK`,
			DevTeam: `Bytedance`,
		},
	},
	{
		Soname: `libShanYCore.so`,
		Zh: SDKLangInfo{
			Label:   `闪验 SDK`,
			DevTeam: `创蓝云智`,
		},
		En: SDKLangInfo{
			Label:   `Flash Verification SDK`,
			DevTeam: `创蓝云智`,
		},
	},
	{
		Soname: `libPixUI_PXPlugin.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmobileglues.so`,
		Zh: SDKLangInfo{
			Label:   `MobileGlues`,
			DevTeam: `MobileGL-Dev`,
		},
		En: SDKLangInfo{
			Label:   `MobileGlues`,
			DevTeam: `MobileGL-Dev`,
		},
	},
	{
		Soname: `libjscexecutor.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libapp.so`,
		Zh: SDKLangInfo{
			Label:   `Flutter`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Flutter`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libshook.so`,
		Zh: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
		En: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
	},
	{
		Soname: `libijm_linker.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libzipw.so`,
		Zh: SDKLangInfo{
			Label:   `zipw`,
			DevTeam: `A dog's life software`,
		},
		En: SDKLangInfo{
			Label:   `zipw`,
			DevTeam: `A dog's life software`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeLinearGradientSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libphotoprocessing.so`,
		Zh: SDKLangInfo{
			Label:   `PhotoProcessing`,
			DevTeam: `Lightbox`,
		},
		En: SDKLangInfo{
			Label:   `PhotoProcessing`,
			DevTeam: `Lightbox`,
		},
	},
	{
		Soname: `libAkSoundEngine.so`,
		Zh: SDKLangInfo{
			Label:   `Wwise`,
			DevTeam: `Audiokinetic`,
		},
		En: SDKLangInfo{
			Label:   `Wwise`,
			DevTeam: `Audiokinetic`,
		},
	},
	{
		Soname: `liblogan.so`,
		Zh: SDKLangInfo{
			Label:   `Logan`,
			DevTeam: `Meituan-Dianping`,
		},
		En: SDKLangInfo{
			Label:   `Logan`,
			DevTeam: `Meituan-Dianping`,
		},
	},
	{
		Soname: `libzxing.so`,
		Zh: SDKLangInfo{
			Label:   `ZXing`,
			DevTeam: `ZXing`,
		},
		En: SDKLangInfo{
			Label:   `ZXing`,
			DevTeam: `ZXing`,
		},
	},
	{
		Soname: `libRTCFFmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libqalmsfboot.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云通信 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Communication SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libijkplayer.so`,
		Zh: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libanddown.so`,
		Zh: SDKLangInfo{
			Label:   `CWAC AndDown`,
			DevTeam: `commonsware`,
		},
		En: SDKLangInfo{
			Label:   `CWAC AndDown`,
			DevTeam: `commonsware`,
		},
	},
	{
		Soname: `libFcitx5Core.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libMNN_Express_Legacy.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libxylinkAIEngineImpOffVulkan.so`,
		Zh: SDKLangInfo{
			Label:   `小鱼易联终端 SDK`,
			DevTeam: `小鱼易联`,
		},
		En: SDKLangInfo{
			Label:   `XYLink Terminal SDK`,
			DevTeam: `小鱼易联`,
		},
	},
	{
		Soname: `libdex2oat10.so`,
		Zh: SDKLangInfo{
			Label:   `dex2oat optimizer`,
			DevTeam: `iamlooper`,
		},
		En: SDKLangInfo{
			Label:   `dex2oat optimizer`,
			DevTeam: `iamlooper`,
		},
	},
	{
		Soname: `libpaddle_mobile_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libandroid-phone-wallet-scanexport.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS 社交分享`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS Social Share`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libkwai-unwind.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libMinilzo.so`,
		Zh: SDKLangInfo{
			Label:   `Minilzo`,
			DevTeam: `Oberhumer`,
		},
		En: SDKLangInfo{
			Label:   `Minilzo`,
			DevTeam: `Oberhumer`,
		},
	},
	{
		Soname: `libmsaoaidsec.so`,
		Zh: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
		En: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
	},
	{
		Soname: `libUnityPlugin.so`,
		Zh: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `libsodium.so`,
		Zh: SDKLangInfo{
			Label:   `Sodium`,
			DevTeam: `jedisct1`,
		},
		En: SDKLangInfo{
			Label:   `Sodium`,
			DevTeam: `jedisct1`,
		},
	},
	{
		Soname: `libkwai-linker.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libCrashSight.so`,
		Zh: SDKLangInfo{
			Label:   `CrashSight`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `CrashSight`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libpaddle_light_api_shared.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libmatrix-artmisc.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblamemp3.so`,
		Zh: SDKLangInfo{
			Label:   `Lame-For-Unity`,
			DevTeam: `3wz`,
		},
		En: SDKLangInfo{
			Label:   `Lame-For-Unity`,
			DevTeam: `3wz`,
		},
	},
	{
		Soname: `libtag.so`,
		Zh: SDKLangInfo{
			Label:   `TagLib`,
			DevTeam: `TagLib`,
		},
		En: SDKLangInfo{
			Label:   `TagLib`,
			DevTeam: `TagLib`,
		},
	},
	{
		Soname: `libfmod.so`,
		Zh: SDKLangInfo{
			Label:   `FMOD`,
			DevTeam: `Firelight`,
		},
		En: SDKLangInfo{
			Label:   `FMOD`,
			DevTeam: `Firelight`,
		},
	},
	{
		Soname: `libandfix_x86.so`,
		Zh: SDKLangInfo{
			Label:   `AndFix`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `AndFix`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `liblib_fb_fbjni.so`,
		Zh: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Yoga`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libmarsstn.so`,
		Zh: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libexpo-modules-core.so`,
		Zh: SDKLangInfo{
			Label:   `Expo`,
			DevTeam: `650 Industries, Inc.`,
		},
		En: SDKLangInfo{
			Label:   `Expo`,
			DevTeam: `650 Industries, Inc.`,
		},
	},
	{
		Soname: `libsecenh_a64.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA Reinforcement`,
			DevTeam: `CFCA`,
		},
		En: SDKLangInfo{
			Label:   `CFCA 加固`,
			DevTeam: `CFCA`,
		},
	},
	{
		Soname: `libTrustUtils.so`,
		Zh: SDKLangInfo{
			Label:   `Trust Wallet Core`,
			DevTeam: `Trustwallet`,
		},
		En: SDKLangInfo{
			Label:   `Trust Wallet Core`,
			DevTeam: `Trustwallet`,
		},
	},
	{
		Soname: `libexecmain.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libzmf.so`,
		Zh: SDKLangInfo{
			Label:   `JusTalk Cloud`,
			DevTeam: `宁波菊风系统软件有限公司`,
		},
		En: SDKLangInfo{
			Label:   `JusTalk Cloud`,
			DevTeam: `宁波菊风系统软件有限公司`,
		},
	},
	{
		Soname: `libsigner.so`,
		Zh: SDKLangInfo{
			Label:   `Adjust SDK`,
			DevTeam: `Adjust`,
		},
		En: SDKLangInfo{
			Label:   `Adjust SDK`,
			DevTeam: `Adjust`,
		},
	},
	{
		Soname: `libyunceng.so`,
		Zh: SDKLangInfo{
			Label:   `阿里云游戏盾`,
			DevTeam: `Aliyun`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Cloud Game Shield`,
			DevTeam: `Aliyun`,
		},
	},
	{
		Soname: `libqcloud_asr_realtime.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云语音识别`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `腾讯云语音识别`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmapscore.so`,
		Zh: SDKLangInfo{
			Label:   `OpenMobileMaps`,
			DevTeam: `openmobilemaps.io`,
		},
		En: SDKLangInfo{
			Label:   `OpenMobileMaps`,
			DevTeam: `openmobilemaps.io`,
		},
	},
	{
		Soname: `libbreakpad.so`,
		Zh: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libapng-drawable.so`,
		Zh: SDKLangInfo{
			Label:   `ApngDrawable`,
			DevTeam: `LINE`,
		},
		En: SDKLangInfo{
			Label:   `ApngDrawable`,
			DevTeam: `LINE`,
		},
	},
	{
		Soname: `libprotobuf.so`,
		Zh: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libimagepipeline.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libonnxruntime.so`,
		Zh: SDKLangInfo{
			Label:   `onnxruntime`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `onnxruntime`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libpinyin.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libhuawei_arengine_impl.so`,
		Zh: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libvad.so`,
		Zh: SDKLangInfo{
			Label:   `vad`,
			DevTeam: `思必驰科技股份有限公司`,
		},
		En: SDKLangInfo{
			Label:   `vad`,
			DevTeam: `思必驰科技股份有限公司`,
		},
	},
	{
		Soname: `liblitr-jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiTr`,
			DevTeam: `Linkedin`,
		},
		En: SDKLangInfo{
			Label:   `LiTr`,
			DevTeam: `Linkedin`,
		},
	},
	{
		Soname: `libTrustWalletCore.so`,
		Zh: SDKLangInfo{
			Label:   `Trust Wallet Core`,
			DevTeam: `Trustwallet`,
		},
		En: SDKLangInfo{
			Label:   `Trust Wallet Core`,
			DevTeam: `Trustwallet`,
		},
	},
	{
		Soname: `libBaiduSpeechSDK.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libapmcrash.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯客户端性能分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Client Performance Analysis`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libquickjs-android-wrapper.so`,
		Zh: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
		En: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
	},
	{
		Soname: `libagora_face_detection_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libibg-native.so`,
		Zh: SDKLangInfo{
			Label:   `Instabug`,
			DevTeam: `Instabug`,
		},
		En: SDKLangInfo{
			Label:   `Instabug`,
			DevTeam: `Instabug`,
		},
	},
	{
		Soname: `libASPEngineCore.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libhdiffpatch.so`,
		Zh: SDKLangInfo{
			Label:   `HDiffPatch`,
			DevTeam: `sisong`,
		},
		En: SDKLangInfo{
			Label:   `HDiffPatch`,
			DevTeam: `sisong`,
		},
	},
	{
		Soname: `libavif-jni.so`,
		Zh: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
		En: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
	},
	{
		Soname: `libakamaibmp.so`,
		Zh: SDKLangInfo{
			Label:   `Akamai Bot Manager`,
			DevTeam: `Akamai`,
		},
		En: SDKLangInfo{
			Label:   `Akamai Bot Manager`,
			DevTeam: `Akamai`,
		},
	},
	{
		Soname: `libchaosvmp.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libagora_video_process_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libreact_codegen_rncore.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libgetuiext3.so`,
		Zh: SDKLangInfo{
			Label:   `个推`,
			DevTeam: `个推`,
		},
		En: SDKLangInfo{
			Label:   `Getui`,
			DevTeam: `个推`,
		},
	},
	{
		Soname: `libopentok.so`,
		Zh: SDKLangInfo{
			Label:   `OpenTok`,
			DevTeam: `TokBox`,
		},
		En: SDKLangInfo{
			Label:   `OpenTok`,
			DevTeam: `TokBox`,
		},
	},
	{
		Soname: `libconscrypt_gmscore_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Conscrypt`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Conscrypt`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmpaas_cppbase.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libnrtc_sdk.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libAPMDump.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libomesStdSco.so`,
		Zh: SDKLangInfo{
			Label:   `OPPO 安全检测 SDK`,
			DevTeam: `OPPO`,
		},
		En: SDKLangInfo{
			Label:   `OPPO 安全检测 SDK`,
			DevTeam: `OPPO`,
		},
	},
	{
		Soname: `libfolly_runtime.so`,
		Zh: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libQmBlur.so`,
		Zh: SDKLangInfo{
			Label:   `QmBlur`,
			DevTeam: `QmDeve`,
		},
		En: SDKLangInfo{
			Label:   `QmBlur`,
			DevTeam: `QmDeve`,
		},
	},
	{
		Soname: `libliquidfun_jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiquidFun`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `LiquidFun`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libreactnativejnifb.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libHKESipCryptor.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
		En: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
	},
	{
		Soname: `libmatrixmrs.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_codegen_safeareacontext.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libcello_native.so`,
		Zh: SDKLangInfo{
			Label:   `Cello`,
			DevTeam: `orangeduck`,
		},
		En: SDKLangInfo{
			Label:   `Cello`,
			DevTeam: `orangeduck`,
		},
	},
	{
		Soname: `libddog.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libspatch.so`,
		Zh: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
		En: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
	},
	{
		Soname: `libcurl.so`,
		Zh: SDKLangInfo{
			Label:   `LibCurl`,
			DevTeam: `curl`,
		},
		En: SDKLangInfo{
			Label:   `LibCurl`,
			DevTeam: `curl`,
		},
	},
	{
		Soname: `libjingle_peerconnection_so.so`,
		Zh: SDKLangInfo{
			Label:   `WebRTC`,
			DevTeam: `WebRTC`,
		},
		En: SDKLangInfo{
			Label:   `WebRTC`,
			DevTeam: `WebRTC`,
		},
	},
	{
		Soname: `libhockey_exception_handler.so`,
		Zh: SDKLangInfo{
			Label:   `App Center`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `App Center`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libopenssl_smime.so`,
		Zh: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
		En: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
	},
	{
		Soname: `libhermes-compile.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libMSDKSystem.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libflutter.so`,
		Zh: SDKLangInfo{
			Label:   `Flutter`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Flutter`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libhiai_ir_infershape.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libvcap_opencl.so`,
		Zh: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libpaddle-mobile.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libsentry.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libhermes-executor-release.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libFirebaseCppMessaging.so`,
		Zh: SDKLangInfo{
			Label:   `Firebase Cloud Messaging`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Firebase Cloud Messaging`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libweexcoreqjs.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsandhook-native.so`,
		Zh: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
		En: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
	},
	{
		Soname: `libmobileffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libweexjss.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libwebpimage.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libGLESv1_CM_angle.so`,
		Zh: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libMegviiDlmk.so`,
		Zh: SDKLangInfo{
			Label:   `人像美化 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Portrait Beautification SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libWCDB.so`,
		Zh: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtable.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libreact_codegen_reactandroidspec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `librxffmpeg-invoke.so`,
		Zh: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
		En: SDKLangInfo{
			Label:   `RxFFmpeg`,
			DevTeam: `microshow`,
		},
	},
	{
		Soname: `libagora-rtc-sdk-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libBreakpadLibrary.so`,
		Zh: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libcfcamkit.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
		En: SDKLangInfo{
			Label:   `CFCA`,
			DevTeam: `中国金融认证中心 (CFCA)`,
		},
	},
	{
		Soname: `libagora_udrm3_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libtxsoundtouch.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云实时音视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Real-time Audio and Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libaudio_preprocessing.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `liblept.so`,
		Zh: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
		En: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
	},
	{
		Soname: `libmibrainvad2.so`,
		Zh: SDKLangInfo{
			Label:   `小爱 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiao Ai SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libarcsoft_videoautozoom.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libpag.so`,
		Zh: SDKLangInfo{
			Label:   `PAG`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PAG`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbaiduprotect_art.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libmp3decoder.so`,
		Zh: SDKLangInfo{
			Label:   `高德地图 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Amap SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libavif.so`,
		Zh: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
		En: SDKLangInfo{
			Label:   `libavif`,
			DevTeam: `Alliance for Open Media`,
		},
	},
	{
		Soname: `libmodpng.so`,
		Zh: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libCGE.so`,
		Zh: SDKLangInfo{
			Label:   `GPUImage-Plus`,
			DevTeam: `wysaid`,
		},
		En: SDKLangInfo{
			Label:   `GPUImage-Plus`,
			DevTeam: `wysaid`,
		},
	},
	{
		Soname: `libdevinfo.so`,
		Zh: SDKLangInfo{
			Label:   `libdevinfo`,
			DevTeam: `Oracle`,
		},
		En: SDKLangInfo{
			Label:   `libdevinfo`,
			DevTeam: `Oracle`,
		},
	},
	{
		Soname: `libYTLiveness.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblibpag.so`,
		Zh: SDKLangInfo{
			Label:   `PAG`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PAG`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libIMEPinyin.so`,
		Zh: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `lib7-Zip-JBinding.so`,
		Zh: SDKLangInfo{
			Label:   `7-Zip`,
			DevTeam: `borisbrodski`,
		},
		En: SDKLangInfo{
			Label:   `7-Zip`,
			DevTeam: `borisbrodski`,
		},
	},
	{
		Soname: `libreact_debug.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libxnnloader.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libpinyinhelper.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libfire.so`,
		Zh: SDKLangInfo{
			Label:   `Sofire`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Sofire`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libalicomphonenumberauthsdk-log-release_alijtca_plus.so`,
		Zh: SDKLangInfo{
			Label:   `号码认证服务`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Number Authentication Service`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libhiai_aar_adapter.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libmibraindec.so`,
		Zh: SDKLangInfo{
			Label:   `小爱 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiao Ai SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.audio.sys.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libapm_astrolabe.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libmmkv.so`,
		Zh: SDKLangInfo{
			Label:   `MMKV`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `MMKV`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libBDSpeechDecoder_V1.so`,
		Zh: SDKLangInfo{
			Label:   `离线语音合成 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Offline Speech Synthesis SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libpolyvffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
		En: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
	},
	{
		Soname: `libstatic-webp.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbaiduprotect_sec_jni_v3_lite_face.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libtermux.so`,
		Zh: SDKLangInfo{
			Label:   `Termux`,
			DevTeam: `Termux`,
		},
		En: SDKLangInfo{
			Label:   `Termux`,
			DevTeam: `Termux`,
		},
	},
	{
		Soname: `libpaddle_lite_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libogg.so`,
		Zh: SDKLangInfo{
			Label:   `ogg`,
			DevTeam: `xiph`,
		},
		En: SDKLangInfo{
			Label:   `ogg`,
			DevTeam: `xiph`,
		},
	},
	{
		Soname: `libreact_codegen_BCLimearGradientSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbucket2-new.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libXD.so`,
		Zh: SDKLangInfo{
			Label:   `XD`,
			DevTeam: `benhardfritz`,
		},
		En: SDKLangInfo{
			Label:   `XD`,
			DevTeam: `benhardfritz`,
		},
	},
	{
		Soname: `libijmDataEncryption_x86_64.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libnms.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libovpnutil.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libzegoliveroom.so`,
		Zh: SDKLangInfo{
			Label:   `即构极速视频 SDK`,
			DevTeam: `即构`,
		},
		En: SDKLangInfo{
			Label:   `Zego Express Video SDK`,
			DevTeam: `即构`,
		},
	},
	{
		Soname: `liblelink.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libBugly-rqd.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libreanimated.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Reanimated`,
			DevTeam: `software-mansion`,
		},
		En: SDKLangInfo{
			Label:   `React Native Reanimated`,
			DevTeam: `software-mansion`,
		},
	},
	{
		Soname: `libAliHALogEngine.so`,
		Zh: SDKLangInfo{
			Label:   `日志服务 SLS`,
			DevTeam: `Aliyun`,
		},
		En: SDKLangInfo{
			Label:   `Log Service SLS`,
			DevTeam: `Aliyun`,
		},
	},
	{
		Soname: `libsentry-android.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libsecuritydevice.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libopusV2JNI.so`,
		Zh: SDKLangInfo{
			Label:   `ExoPlayer`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ExoPlayer`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libquack.so`,
		Zh: SDKLangInfo{
			Label:   `Quack`,
			DevTeam: `koush`,
		},
		En: SDKLangInfo{
			Label:   `Quack`,
			DevTeam: `koush`,
		},
	},
	{
		Soname: `libreact_config.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libwechatbacktrace.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbd_license.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libvolcenginertc.so`,
		Zh: SDKLangInfo{
			Label:   `实时音视频`,
			DevTeam: `火山引擎`,
		},
		En: SDKLangInfo{
			Label:   `Real-time Audio and Video`,
			DevTeam: `火山引擎`,
		},
	},
	{
		Soname: `librime.so`,
		Zh: SDKLangInfo{
			Label:   `Rime Core`,
			DevTeam: `Rime (create by lotem)`,
		},
		En: SDKLangInfo{
			Label:   `Rime Core`,
			DevTeam: `Rime (create by lotem)`,
		},
	},
	{
		Soname: `libweexcore.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libcrashpad_handler.so`,
		Zh: SDKLangInfo{
			Label:   `crashpad`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `crashpad`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `libhiai.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libmsc.so`,
		Zh: SDKLangInfo{
			Label:   `讯飞 SDK`,
			DevTeam: `科大讯飞`,
		},
		En: SDKLangInfo{
			Label:   `iFlytek SDK`,
			DevTeam: `科大讯飞`,
		},
	},
	{
		Soname: `libbytevc0.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libads-c.so`,
		Zh: SDKLangInfo{
			Label:   `OPPO 广告 SDK`,
			DevTeam: `OPPO`,
		},
		En: SDKLangInfo{
			Label:   `OPPO 广告 SDK`,
			DevTeam: `OPPO`,
		},
	},
	{
		Soname: `libagora_super_resolution_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libdeflate.so`,
		Zh: SDKLangInfo{
			Label:   `libdeflate`,
			DevTeam: `ebiggers`,
		},
		En: SDKLangInfo{
			Label:   `libdeflate`,
			DevTeam: `ebiggers`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeLinerGradientSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libhandwriting.so`,
		Zh: SDKLangInfo{
			Label:   `即构极速视频 SDK`,
			DevTeam: `即构`,
		},
		En: SDKLangInfo{
			Label:   `Zego Express Video SDK`,
			DevTeam: `即构`,
		},
	},
	{
		Soname: `libmla.so`,
		Zh: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
		En: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
	},
	{
		Soname: `libed25519.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libSecShell-x86.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libalicomphonenumberauthsdk-release_alijtca_plus.so`,
		Zh: SDKLangInfo{
			Label:   `号码认证服务`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Number Authentication Service`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libagora_video_encoder_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libcrashpad_client.so`,
		Zh: SDKLangInfo{
			Label:   `crashpad`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `crashpad`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `liblanguage_id_l2c_jni.so`,
		Zh: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libomp.so`,
		Zh: SDKLangInfo{
			Label:   `LibOMP`,
			DevTeam: `OMPI-X`,
		},
		En: SDKLangInfo{
			Label:   `LibOMP`,
			DevTeam: `OMPI-X`,
		},
	},
	{
		Soname: `libsgcore.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains_v2.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libreact_codegen_reactnativekeyboardcontroller.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-keyboard-controller`,
			DevTeam: `kirillzyusko`,
		},
		En: SDKLangInfo{
			Label:   `react-native-keyboard-controller`,
			DevTeam: `kirillzyusko`,
		},
	},
	{
		Soname: `libHttpDnsPlugin.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云 HTTPDNS`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud HTTPDNS`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libfuai.so`,
		Zh: SDKLangInfo{
			Label:   `Faceunity Nama SDK`,
			DevTeam: `杭州相芯科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Faceunity Nama SDK`,
			DevTeam: `杭州相芯科技有限公司`,
		},
	},
	{
		Soname: `libreact-native-mmkv.so`,
		Zh: SDKLangInfo{
			Label:   `React Native MMKV`,
			DevTeam: `mrousavy`,
		},
		En: SDKLangInfo{
			Label:   `React Native MMKV`,
			DevTeam: `mrousavy`,
		},
	},
	{
		Soname: `liblwjgl_opengl.so`,
		Zh: SDKLangInfo{
			Label:   `LWJGL`,
			DevTeam: `LWJGL`,
		},
		En: SDKLangInfo{
			Label:   `LWJGL`,
			DevTeam: `LWJGL`,
		},
	},
	{
		Soname: `libFFmpeg_video.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libtvm4j_runtime_packed.so`,
		Zh: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
		En: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
	},
	{
		Soname: `liblbffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libweexjst.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsentry-fd-monitor.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libcovault-appsec.so`,
		Zh: SDKLangInfo{
			Label:   `AppSealing`,
			DevTeam: `INKA Entworks`,
		},
		En: SDKLangInfo{
			Label:   `AppSealing`,
			DevTeam: `INKA Entworks`,
		},
	},
	{
		Soname: `libopusJNI.so`,
		Zh: SDKLangInfo{
			Label:   `opus-jni`,
			DevTeam: `LabyMod`,
		},
		En: SDKLangInfo{
			Label:   `opus-jni`,
			DevTeam: `LabyMod`,
		},
	},
	{
		Soname: `libqimei.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔终端 ID 精准识别体系`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `QIMEI SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblynx_v8_bridge.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libMMASignature.so`,
		Zh: SDKLangInfo{
			Label:   `中国无线营销联盟通用 SDK`,
			DevTeam: `中国无线营销联盟`,
		},
		En: SDKLangInfo{
			Label:   `China Wireless Marketing Alliance General SDK`,
			DevTeam: `中国无线营销联盟`,
		},
	},
	{
		Soname: `libYTKeyboard.so`,
		Zh: SDKLangInfo{
			Label:   `屹通 SDK`,
			DevTeam: `上海屹通信息科技发展有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Yitong SDK`,
			DevTeam: `上海屹通信息科技发展有限公司`,
		},
	},
	{
		Soname: `libxlua.so`,
		Zh: SDKLangInfo{
			Label:   `xLua`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `xLua`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblynx.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libhiai_ir.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libYTImageRefiner.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `librninstance.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libreact_codegen_rnpicker.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-picker`,
			DevTeam: `software-mansion`,
		},
		En: SDKLangInfo{
			Label:   `react-native-picker`,
			DevTeam: `software-mansion`,
		},
	},
	{
		Soname: `libavutil.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libmarsnative.so`,
		Zh: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblayaair.so`,
		Zh: SDKLangInfo{
			Label:   `LayaAir`,
			DevTeam: `Layabox`,
		},
		En: SDKLangInfo{
			Label:   `LayaAir`,
			DevTeam: `Layabox`,
		},
	},
	{
		Soname: `libzbarjni.so`,
		Zh: SDKLangInfo{
			Label:   `ZBar`,
			DevTeam: `spadix`,
		},
		En: SDKLangInfo{
			Label:   `ZBar`,
			DevTeam: `spadix`,
		},
	},
	{
		Soname: `libGVoice.so`,
		Zh: SDKLangInfo{
			Label:   `游戏语音 GVoice`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Game Voice GVoice`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libface_detector_v2_jni.so`,
		Zh: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libfb.so`,
		Zh: SDKLangInfo{
			Label:   `Facebook SDK`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Facebook SDK`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libvideodec.so`,
		Zh: SDKLangInfo{
			Label:   `videodec`,
			DevTeam: `Parrot-Developers`,
		},
		En: SDKLangInfo{
			Label:   `videodec`,
			DevTeam: `Parrot-Developers`,
		},
	},
	{
		Soname: `libCrashSightPlugin.so`,
		Zh: SDKLangInfo{
			Label:   `CrashSight`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `CrashSight`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libpl_droidsonroids_gif.so`,
		Zh: SDKLangInfo{
			Label:   `android-gif-drawable`,
			DevTeam: `koral--`,
		},
		En: SDKLangInfo{
			Label:   `android-gif-drawable`,
			DevTeam: `koral--`,
		},
	},
	{
		Soname: `libtencentloc.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libxhook_lib.so`,
		Zh: SDKLangInfo{
			Label:   `xHook`,
			DevTeam: `iQIYI`,
		},
		En: SDKLangInfo{
			Label:   `xHook`,
			DevTeam: `iQIYI`,
		},
	},
	{
		Soname: `libYTFaceReflect.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libDexHelper-x86_64.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libvcaplite_android.so`,
		Zh: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libavformat.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libagora_spatial_audio_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libuptsmaddon.so`,
		Zh: SDKLangInfo{
			Label:   `银联 SDK`,
			DevTeam: `银联`,
		},
		En: SDKLangInfo{
			Label:   `UnionPay SDK`,
			DevTeam: `银联`,
		},
	},
	{
		Soname: `libgio.so`,
		Zh: SDKLangInfo{
			Label:   `Gio`,
			DevTeam: `Gio`,
		},
		En: SDKLangInfo{
			Label:   `Gio`,
			DevTeam: `Gio`,
		},
	},
	{
		Soname: `libBugly.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libcalculator_domains_skel.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libpolyvplayer.so`,
		Zh: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
		En: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
	},
	{
		Soname: `libCrasheyeNDK.so`,
		Zh: SDKLangInfo{
			Label:   `Crasheye`,
			DevTeam: `西山居`,
		},
		En: SDKLangInfo{
			Label:   `Crasheye`,
			DevTeam: `西山居`,
		},
	},
	{
		Soname: `libmono-native.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libYTFaceTracker.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libagora_mpg123.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libfuse_jni.so`,
		Zh: SDKLangInfo{
			Label:   `libfuse`,
			DevTeam: `libfuse`,
		},
		En: SDKLangInfo{
			Label:   `libfuse`,
			DevTeam: `libfuse`,
		},
	},
	{
		Soname: `libdexjni.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libagora_jnd_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libbrotlidec.so`,
		Zh: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libKF5ConfigCore_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libwechatxlog.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmatrix-signaledhook.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libwoff2_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Google WOFF2`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google WOFF2`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libfaceidlivenessv5.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libaikl_calc_arm.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `liboctopus.so`,
		Zh: SDKLangInfo{
			Label:   `liboctopus`,
			DevTeam: `opendesigndev`,
		},
		En: SDKLangInfo{
			Label:   `liboctopus`,
			DevTeam: `opendesigndev`,
		},
	},
	{
		Soname: `libRDTAPIs.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libhippy.so`,
		Zh: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libvcap_npu.so`,
		Zh: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libYTPoseDetect.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libgdx.so`,
		Zh: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
		En: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
	},
	{
		Soname: `libunicode.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libKF5Archive_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libAPMlog.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libgcloud.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 GCloud`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Games GCloud`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libhs.so`,
		Zh: SDKLangInfo{
			Label:   `Sofire`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Sofire`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libtapsdkcore.so`,
		Zh: SDKLangInfo{
			Label:   `TapSDK`,
			DevTeam: `TapTap`,
		},
		En: SDKLangInfo{
			Label:   `TapSDK`,
			DevTeam: `TapTap`,
		},
	},
	{
		Soname: `libcpucl.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libGLESv2_angle.so`,
		Zh: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libscannative.so`,
		Zh: SDKLangInfo{
			Label:   `HMS Scan Kit`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `HMS Scan Kit`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libarcsoft_panorama_burstcapture.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libepic.so`,
		Zh: SDKLangInfo{
			Label:   `Epic`,
			DevTeam: `weishu`,
		},
		En: SDKLangInfo{
			Label:   `Epic`,
			DevTeam: `weishu`,
		},
	},
	{
		Soname: `libperf.so`,
		Zh: SDKLangInfo{
			Label:   `perf`,
			DevTeam: `theonewolf`,
		},
		En: SDKLangInfo{
			Label:   `perf`,
			DevTeam: `theonewolf`,
		},
	},
	{
		Soname: `libHiAI_dl.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `liblog4a-lib.so`,
		Zh: SDKLangInfo{
			Label:   `Log4a`,
			DevTeam: `pqpo`,
		},
		En: SDKLangInfo{
			Label:   `Log4a`,
			DevTeam: `pqpo`,
		},
	},
	{
		Soname: `libijm-emulator.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libsangforsdk.so`,
		Zh: SDKLangInfo{
			Label:   `Sangfor`,
			DevTeam: `深信服`,
		},
		En: SDKLangInfo{
			Label:   `Sangfor`,
			DevTeam: `深信服`,
		},
	},
	{
		Soname: `libxyvp.so`,
		Zh: SDKLangInfo{
			Label:   `小鱼易联终端 SDK`,
			DevTeam: `小鱼易联`,
		},
		En: SDKLangInfo{
			Label:   `XYLink Terminal SDK`,
			DevTeam: `小鱼易联`,
		},
	},
	{
		Soname: `libmiio.so`,
		Zh: SDKLangInfo{
			Label:   `小米 IoT SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiaomi IoT SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libijksdl.so`,
		Zh: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libnesec_x86.so`,
		Zh: SDKLangInfo{
			Label:   `网易易盾`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Yidun`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libr_upgrade_lib.so`,
		Zh: SDKLangInfo{
			Label:   `RUpgrade`,
			DevTeam: `rhymelph`,
		},
		En: SDKLangInfo{
			Label:   `RUpgrade`,
			DevTeam: `rhymelph`,
		},
	},
	{
		Soname: `libAgoraMediaPlayer.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libmonodroid_bundle_app.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libsqlite-jni.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libkoom-java.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libfluidsynth.so`,
		Zh: SDKLangInfo{
			Label:   `FluidSynth`,
			DevTeam: `FluidSynth`,
		},
		En: SDKLangInfo{
			Label:   `FluidSynth`,
			DevTeam: `FluidSynth`,
		},
	},
	{
		Soname: `libmatrix-hookcommon.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libvad.dnn.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libhttpdns.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云 HTTPDNS`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud HTTPDNS`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_codegen_KmReactEntrySpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libadblockplus-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Adblock Android SDK`,
			DevTeam: `adblockplus`,
		},
		En: SDKLangInfo{
			Label:   `Adblock Android SDK`,
			DevTeam: `adblockplus`,
		},
	},
	{
		Soname: `libxmcore.so`,
		Zh: SDKLangInfo{
			Label:   `xmCore`,
			DevTeam: `luxiaoming`,
		},
		En: SDKLangInfo{
			Label:   `xmCore`,
			DevTeam: `luxiaoming`,
		},
	},
	{
		Soname: `libluaaddonloader.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 Lua`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Lua`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libIOTCAPIsT.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libMGBeauty.so`,
		Zh: SDKLangInfo{
			Label:   `人像美化 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Portrait Beautification SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libjpeg.so`,
		Zh: SDKLangInfo{
			Label:   `LibJPEG`,
			DevTeam: `Independent JPEG Group`,
		},
		En: SDKLangInfo{
			Label:   `LibJPEG`,
			DevTeam: `Independent JPEG Group`,
		},
	},
	{
		Soname: `libTDataMaster.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtersafe2.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云手游安全`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Mobile Game Security`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libhermes-executor-debug.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libBurstLinker.so`,
		Zh: SDKLangInfo{
			Label:   `BurstLinker`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `BurstLinker`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libhippycanvas.so`,
		Zh: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libwind.so`,
		Zh: SDKLangInfo{
			Label:   `微博 SDK`,
			DevTeam: `Weibo`,
		},
		En: SDKLangInfo{
			Label:   `Weibo SDK`,
			DevTeam: `Weibo`,
		},
	},
	{
		Soname: `libopencv_core.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `libmatrix-traffic.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libc++_shared.so`,
		Zh: SDKLangInfo{
			Label:   `C++ 共享库`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `C++ Shared Library`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libreact_codegen_RNCWebViewSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libxnn_invoker.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libpgw.so`,
		Zh: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
		En: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
	},
	{
		Soname: `libcjson.so`,
		Zh: SDKLangInfo{
			Label:   `cJSON`,
			DevTeam: `cJSON`,
		},
		En: SDKLangInfo{
			Label:   `cJSON`,
			DevTeam: `cJSON`,
		},
	},
	{
		Soname: `libpatcher.so`,
		Zh: SDKLangInfo{
			Label:   `AndroidUtils`,
			DevTeam: `zhengxiaopeng`,
		},
		En: SDKLangInfo{
			Label:   `AndroidUtils`,
			DevTeam: `zhengxiaopeng`,
		},
	},
	{
		Soname: `libvo2.so`,
		Zh: SDKLangInfo{
			Label:   `libmpeg2`,
			DevTeam: `Sam Hocevar`,
		},
		En: SDKLangInfo{
			Label:   `libmpeg2`,
			DevTeam: `Sam Hocevar`,
		},
	},
	{
		Soname: `libtvm4j.so`,
		Zh: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
		En: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.kws.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libandroidfrontend.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libSnpeGpu.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libYTImageRefinerPub.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libsecexe.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libsqlc-native-driver.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-sqlite-storage`,
			DevTeam: `Andrzej Porebski`,
		},
		En: SDKLangInfo{
			Label:   `react-native-sqlite-storage`,
			DevTeam: `Andrzej Porebski`,
		},
	},
	{
		Soname: `libhsynz.so`,
		Zh: SDKLangInfo{
			Label:   `hsynz`,
			DevTeam: `sisong`,
		},
		En: SDKLangInfo{
			Label:   `hsynz`,
			DevTeam: `sisong`,
		},
	},
	{
		Soname: `libreact_devsupportjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libBeaconDT.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔 DataTalk`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Beacon DataTalk`,
			DevTeam: `Tencent Beacon`,
		},
	},
	{
		Soname: `libmpaas_common_teclabilities.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libfoza.so`,
		Zh: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
		En: SDKLangInfo{
			Label:   `SHook`,
			DevTeam: `DIE.LU`,
		},
	},
	{
		Soname: `libhuawei_arengine_jni.so`,
		Zh: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libHttpDnsCore.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云 HTTPDNS`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud HTTPDNS`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libxyvodsdk.so`,
		Zh: SDKLangInfo{
			Label:   `小鱼易联终端 SDK`,
			DevTeam: `小鱼易联`,
		},
		En: SDKLangInfo{
			Label:   `XYLink Terminal SDK`,
			DevTeam: `小鱼易联`,
		},
	},
	{
		Soname: `libBaiduSpeechAEC.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libBeacon.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯流量联盟`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Traffic Alliance`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_codegen_RNMmkvSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libquickjs_jni.so`,
		Zh: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
		En: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
	},
	{
		Soname: `libbaiduprotect_dump.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libreact_codegen_qrn_adr.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libclipboard.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libcocos2dcpp.so`,
		Zh: SDKLangInfo{
			Label:   `cocos2d-cpp`,
			DevTeam: `cocos2d-cpp`,
		},
		En: SDKLangInfo{
			Label:   `cocos2d-cpp`,
			DevTeam: `cocos2d-cpp`,
		},
	},
	{
		Soname: `libeffect_proxy.so`,
		Zh: SDKLangInfo{
			Label:   `CV SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `CV SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libeffect.so`,
		Zh: SDKLangInfo{
			Label:   `CV SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `CV SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libdexvmp.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `liblynx_krypton.so`,
		Zh: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Lynx`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libhpplayvideo19.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libhcl.so`,
		Zh: SDKLangInfo{
			Label:   `HCL2`,
			DevTeam: `Nukosuke`,
		},
		En: SDKLangInfo{
			Label:   `HCL2`,
			DevTeam: `Nukosuke`,
		},
	},
	{
		Soname: `libvlc.so`,
		Zh: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
		En: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
	},
	{
		Soname: `libP2PTunnelAPIsT.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libmegazord.so`,
		Zh: SDKLangInfo{
			Label:   `Megazord`,
			DevTeam: `Mozilla`,
		},
		En: SDKLangInfo{
			Label:   `Megazord`,
			DevTeam: `Mozilla`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libvivo_account_sdk.so`,
		Zh: SDKLangInfo{
			Label:   `vivo 网游联运 SDK`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `vivo Online Game Co-operation SDK`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libsentry-gwp-asan.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libcroppy.so`,
		Zh: SDKLangInfo{
			Label:   `croppy`,
			DevTeam: `kekland`,
		},
		En: SDKLangInfo{
			Label:   `croppy`,
			DevTeam: `kekland`,
		},
	},
	{
		Soname: `libcrashlytics-handler.so`,
		Zh: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libKF5Completion_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libGXAnalyzeAndroid.so`,
		Zh: SDKLangInfo{
			Label:   `GaiaX`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `GaiaX`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `librnmmkv.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-mmkv-storage`,
			DevTeam: `Ammar Ahmed`,
		},
		En: SDKLangInfo{
			Label:   `react-native-mmkv-storage`,
			DevTeam: `Ammar Ahmed`,
		},
	},
	{
		Soname: `libOvenPlayerMCVideo.so`,
		Zh: SDKLangInfo{
			Label:   `OvenPlayer`,
			DevTeam: `AirenSoft`,
		},
		En: SDKLangInfo{
			Label:   `OvenPlayer`,
			DevTeam: `AirenSoft`,
		},
	},
	{
		Soname: `libmatrix-mallctl.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmediainfo.so`,
		Zh: SDKLangInfo{
			Label:   `MediaInfoLib`,
			DevTeam: `MediaArea`,
		},
		En: SDKLangInfo{
			Label:   `MediaInfoLib`,
			DevTeam: `MediaArea`,
		},
	},
	{
		Soname: `libijkffext.so`,
		Zh: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
		En: SDKLangInfo{
			Label:   `IJKPlayer`,
			DevTeam: `Bilibili`,
		},
	},
	{
		Soname: `libuptsmaddonmi.so`,
		Zh: SDKLangInfo{
			Label:   `银联 SDK`,
			DevTeam: `银联`,
		},
		En: SDKLangInfo{
			Label:   `UnionPay SDK`,
			DevTeam: `银联`,
		},
	},
	{
		Soname: `libxluau.so`,
		Zh: SDKLangInfo{
			Label:   `xLua`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `xLua`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtxsdl.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云短视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Short Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libgvrbase.so`,
		Zh: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libsgsgmiddletier.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libvcap.so`,
		Zh: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libbson.so`,
		Zh: SDKLangInfo{
			Label:   `MongoDB`,
			DevTeam: `MongoDB`,
		},
		En: SDKLangInfo{
			Label:   `MongoDB`,
			DevTeam: `MongoDB`,
		},
	},
	{
		Soname: `libbugsnag-root-detection.so`,
		Zh: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
		En: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
	},
	{
		Soname: `libkoom-thread.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libtbs.resources.pak.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmediapipe_tasks_vision_jni.so`,
		Zh: SDKLangInfo{
			Label:   `MediaPipe Tasks`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `MediaPipe Tasks`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libMegActionFmpJni.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libhuawei_arengine.so`,
		Zh: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libreact_codegen_ReactEntrySpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbitmaps.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libquiche.so`,
		Zh: SDKLangInfo{
			Label:   `quiche`,
			DevTeam: `Cloudflare`,
		},
		En: SDKLangInfo{
			Label:   `quiche`,
			DevTeam: `Cloudflare`,
		},
	},
	{
		Soname: `libjsengine-loadso.so`,
		Zh: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
		En: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
	},
	{
		Soname: `libmediautil_v6.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `libJavaScriptCore.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.kws.ort.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libhyphenate.so`,
		Zh: SDKLangInfo{
			Label:   `环信 IM`,
			DevTeam: `环信`,
		},
		En: SDKLangInfo{
			Label:   `Easemob IM`,
			DevTeam: `环信`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains_system.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.java.bindings.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libgojni.so`,
		Zh: SDKLangInfo{
			Label:   `Golang`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Golang`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libfilament-utils-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmindspore-lite.so`,
		Zh: SDKLangInfo{
			Label:   `MindSpore Lite`,
			DevTeam: `Huawei, MindSpore Community`,
		},
		En: SDKLangInfo{
			Label:   `MindSpore Lite`,
			DevTeam: `Huawei, MindSpore Community`,
		},
	},
	{
		Soname: `libmatrix-junwind.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libjserrorhandler.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `liblocationbeaconid.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libAisound.so`,
		Zh: SDKLangInfo{
			Label:   `离线语音合成`,
			DevTeam: `iFlyTek`,
		},
		En: SDKLangInfo{
			Label:   `Offline Speech Synthesis`,
			DevTeam: `iFlyTek`,
		},
	},
	{
		Soname: `libgpg.so`,
		Zh: SDKLangInfo{
			Label:   `Google Play Games plugin for Unity`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google Play Games plugin for Unity`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libtiff.so`,
		Zh: SDKLangInfo{
			Label:   `Android-TiffBitmapFactory`,
			DevTeam: `Beyka`,
		},
		En: SDKLangInfo{
			Label:   `Android-TiffBitmapFactory`,
			DevTeam: `Beyka`,
		},
	},
	{
		Soname: `libunity.so`,
		Zh: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `liblkjingle_peerconnection_so.so`,
		Zh: SDKLangInfo{
			Label:   `WebRTC`,
			DevTeam: `WebRTC`,
		},
		En: SDKLangInfo{
			Label:   `WebRTC`,
			DevTeam: `WebRTC`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeWidgetSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libAudioProc.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libsgsecuritybody.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libWXVoice.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云语音识别`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `腾讯云语音识别`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libavcodec.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libvolcenginertcaudio.so`,
		Zh: SDKLangInfo{
			Label:   `实时音视频`,
			DevTeam: `火山引擎`,
		},
		En: SDKLangInfo{
			Label:   `Real-time Audio and Video`,
			DevTeam: `火山引擎`,
		},
	},
	{
		Soname: `libyoyo.so`,
		Zh: SDKLangInfo{
			Label:   `GameMaker Studio`,
			DevTeam: `YoYo Games`,
		},
		En: SDKLangInfo{
			Label:   `GameMaker Studio`,
			DevTeam: `YoYo Games`,
		},
	},
	{
		Soname: `libFcitx5Config.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libsqlitejdbc.so`,
		Zh: SDKLangInfo{
			Label:   `SQlite JDBC`,
			DevTeam: `xerial`,
		},
		En: SDKLangInfo{
			Label:   `SQlite JDBC`,
			DevTeam: `xerial`,
		},
	},
	{
		Soname: `libijmDataEncryption.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libarcore_c.so`,
		Zh: SDKLangInfo{
			Label:   `ARCore`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ARCore`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libagora-core.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libisar.so`,
		Zh: SDKLangInfo{
			Label:   `Isar Database`,
			DevTeam: `isar`,
		},
		En: SDKLangInfo{
			Label:   `Isar Database`,
			DevTeam: `isar`,
		},
	},
	{
		Soname: `libtraeimp-rtmp.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云实时音视频`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Real-time Audio and Video`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libexecmain_x86.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libpairipcore.so`,
		Zh: SDKLangInfo{
			Label:   `PairIP`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `PairIP`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmnncv.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libPixVideo.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libswan-native-lib.so`,
		Zh: SDKLangInfo{
			Label:   `智能小程序`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Smart Mini Program`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `librsjni.so`,
		Zh: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libcommon_lib_uc.so`,
		Zh: SDKLangInfo{
			Label:   `UC 内核`,
			DevTeam: `UC`,
		},
		En: SDKLangInfo{
			Label:   `UC 内核`,
			DevTeam: `UC`,
		},
	},
	{
		Soname: `libwechatcommon.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libvpxV2JNI.so`,
		Zh: SDKLangInfo{
			Label:   `WebM VP8/VP9`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `WebM VP8/VP9`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnss3.so`,
		Zh: SDKLangInfo{
			Label:   `Mozilla NSS`,
			DevTeam: `Mozilla`,
		},
		En: SDKLangInfo{
			Label:   `Mozilla NSS`,
			DevTeam: `Mozilla`,
		},
	},
	{
		Soname: `libwukong_ua.so`,
		Zh: SDKLangInfo{
			Label:   `Wukong`,
			DevTeam: `哈啰`,
		},
		En: SDKLangInfo{
			Label:   `Wukong`,
			DevTeam: `哈啰`,
		},
	},
	{
		Soname: `libgdx-freetype.so`,
		Zh: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
		En: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
	},
	{
		Soname: `libagora_video_quality_analyzer_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libpano_video_renderer.so`,
		Zh: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbarhopper_v3.so`,
		Zh: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libkoom-fast-dump.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libexecutorch.so`,
		Zh: SDKLangInfo{
			Label:   `Executorch`,
			DevTeam: `Facebook/Pytorch`,
		},
		En: SDKLangInfo{
			Label:   `Executorch`,
			DevTeam: `Facebook/Pytorch`,
		},
	},
	{
		Soname: `libNetHTProtect.so`,
		Zh: SDKLangInfo{
			Label:   `网易易盾智能反外挂`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Easy Shield Intelligent Anti-Cheat`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libcrashlytics-trampoline.so`,
		Zh: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libyx_alog.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libagora_drm_loader_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libpixulCurl.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libMsdkAdapter.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libssl_chaquopy.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libKF5CoreAddons_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libne_audio.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libapmdalvik.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯客户端性能分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Client Performance Analysis`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libagora-rtc-sdk.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libexpo-sqlite.so`,
		Zh: SDKLangInfo{
			Label:   `Expo SQLite`,
			DevTeam: `Expo`,
		},
		En: SDKLangInfo{
			Label:   `Expo SQLite`,
			DevTeam: `Expo`,
		},
	},
	{
		Soname: `libsecurityenv.so`,
		Zh: SDKLangInfo{
			Label:   `U-Game 金融风控组件`,
			DevTeam: `Umeng`,
		},
		En: SDKLangInfo{
			Label:   `U-Game Financial Risk Control Component`,
			DevTeam: `Umeng`,
		},
	},
	{
		Soname: `libleboAudio.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libhiai_debugger.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `liblbplayer.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libyxbase.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libaikl_cluster_arm.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libarkui_android.so`,
		Zh: SDKLangInfo{
			Label:   `ArkUI-X`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `ArkUI-X`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libtensorflowlite_jni_stable.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libflexbox.so`,
		Zh: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libneliveplayer.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libProtobufLite.so`,
		Zh: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libagora_ai_noise_suppression_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libcalculator_skel.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libaiui.so`,
		Zh: SDKLangInfo{
			Label:   `AIUI`,
			DevTeam: `iFlyTek`,
		},
		En: SDKLangInfo{
			Label:   `AIUI`,
			DevTeam: `iFlyTek`,
		},
	},
	{
		Soname: `libargon2-wrapper-0.5.so`,
		Zh: SDKLangInfo{
			Label:   `Argon2`,
			DevTeam: `skynet-im`,
		},
		En: SDKLangInfo{
			Label:   `Argon2`,
			DevTeam: `skynet-im`,
		},
	},
	{
		Soname: `libtxmapvis.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libflutter_pty.so`,
		Zh: SDKLangInfo{
			Label:   `Flutter Pty`,
			DevTeam: `TerminalStudio`,
		},
		En: SDKLangInfo{
			Label:   `Flutter Pty`,
			DevTeam: `TerminalStudio`,
		},
	},
	{
		Soname: `libMonoPosixHelper.so`,
		Zh: SDKLangInfo{
			Label:   `Unity Mono`,
			DevTeam: `Unity`,
		},
		En: SDKLangInfo{
			Label:   `Unity Mono`,
			DevTeam: `Unity`,
		},
	},
	{
		Soname: `libhermesinstancejni.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libimselector.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libfastcv.so`,
		Zh: SDKLangInfo{
			Label:   `FastCV`,
			DevTeam: `Qualcomm`,
		},
		En: SDKLangInfo{
			Label:   `FastCV`,
			DevTeam: `Qualcomm`,
		},
	},
	{
		Soname: `libBugly-yaq.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libagora-soundtouch.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libpngt.so`,
		Zh: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
		En: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
	},
	{
		Soname: `libczxing.so`,
		Zh: SDKLangInfo{
			Label:   `CZXing`,
			DevTeam: `devilsen`,
		},
		En: SDKLangInfo{
			Label:   `CZXing`,
			DevTeam: `devilsen`,
		},
	},
	{
		Soname: `libgnustl_shared.so`,
		Zh: SDKLangInfo{
			Label:   `C++ 共享库`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `C++ Shared Library`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libtvm4j_runtime.so`,
		Zh: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
		En: SDKLangInfo{
			Label:   `TVM`,
			DevTeam: `Apache`,
		},
	},
	{
		Soname: `libmatrix-memoryhook.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblbwebpdec.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libarcsoft_beauty_ex.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libpcre2.so`,
		Zh: SDKLangInfo{
			Label:   `pcre2`,
			DevTeam: `PCRE2Project`,
		},
		En: SDKLangInfo{
			Label:   `pcre2`,
			DevTeam: `PCRE2Project`,
		},
	},
	{
		Soname: `libeasymobile.so`,
		Zh: SDKLangInfo{
			Label:   `Easy Mobile`,
			DevTeam: `SgLib Games`,
		},
		En: SDKLangInfo{
			Label:   `Easy Mobile`,
			DevTeam: `SgLib Games`,
		},
	},
	{
		Soname: `librsjni_androidx.so`,
		Zh: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libmpaas_teclability_interface.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libYTLipReader.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbaiduprotect_sec_jni.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libreact-package-registry.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libWeexEagle.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libfmodL.so`,
		Zh: SDKLangInfo{
			Label:   `FMOD`,
			DevTeam: `Firelight`,
		},
		En: SDKLangInfo{
			Label:   `FMOD`,
			DevTeam: `Firelight`,
		},
	},
	{
		Soname: `libalicomphonenumberauthsdk_core.so`,
		Zh: SDKLangInfo{
			Label:   `号码认证服务`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Number Authentication Service`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsnapdragon_services_qape.qti.so`,
		Zh: SDKLangInfo{
			Label:   `QESDK`,
			DevTeam: `Qualcomm`,
		},
		En: SDKLangInfo{
			Label:   `QESDK`,
			DevTeam: `Qualcomm`,
		},
	},
	{
		Soname: `libmpaas_network_kotlin_adapter.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libgocryptfs_jni.so`,
		Zh: SDKLangInfo{
			Label:   `gocryptsfs`,
			DevTeam: `@rfjakob, etc.`,
		},
		En: SDKLangInfo{
			Label:   `gocryptsfs`,
			DevTeam: `@rfjakob, etc.`,
		},
	},
	{
		Soname: `libtracepath.so`,
		Zh: SDKLangInfo{
			Label:   `tracepath`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `tracepath`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libweibosdkcore.so`,
		Zh: SDKLangInfo{
			Label:   `微博 SDK`,
			DevTeam: `Weibo`,
		},
		En: SDKLangInfo{
			Label:   `Weibo SDK`,
			DevTeam: `Weibo`,
		},
	},
	{
		Soname: `libartemis.so`,
		Zh: SDKLangInfo{
			Label:   `Artemis Engine`,
			DevTeam: `御影 (Mikage)`,
		},
		En: SDKLangInfo{
			Label:   `Artemis Engine`,
			DevTeam: `御影 (Mikage)`,
		},
	},
	{
		Soname: `libmediashow.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `libgocryptfs.so`,
		Zh: SDKLangInfo{
			Label:   `gocryptsfs`,
			DevTeam: `@rfjakob, etc.`,
		},
		En: SDKLangInfo{
			Label:   `gocryptsfs`,
			DevTeam: `@rfjakob, etc.`,
		},
	},
	{
		Soname: `liblouiswrap.so`,
		Zh: SDKLangInfo{
			Label:   `Liblouis`,
			DevTeam: `Liblouis`,
		},
		En: SDKLangInfo{
			Label:   `Liblouis`,
			DevTeam: `Liblouis`,
		},
	},
	{
		Soname: `libarcsoft_dualcam_refocus_gallery.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libopenal.so`,
		Zh: SDKLangInfo{
			Label:   `OpenAL`,
			DevTeam: `Ubuntu MOTU Developers`,
		},
		En: SDKLangInfo{
			Label:   `OpenAL`,
			DevTeam: `Ubuntu MOTU Developers`,
		},
	},
	{
		Soname: `libarcore_sdk_jni.so`,
		Zh: SDKLangInfo{
			Label:   `ARCore`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ARCore`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libkrisp-audio-sdk.so`,
		Zh: SDKLangInfo{
			Label:   `Krisp AI 语音 SDK`,
			DevTeam: `krisp`,
		},
		En: SDKLangInfo{
			Label:   `Krisp AI Voice SDK`,
			DevTeam: `krisp`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.core.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libgloudcore.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 GCloud`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Games GCloud`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbaiduprotect_x86.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libsentry-hook.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libsdk_patcher_jni.so`,
		Zh: SDKLangInfo{
			Label:   `小米应用检查更新 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiaomi App Update Check SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libvideokit_hdrvivid.so`,
		Zh: SDKLangInfo{
			Label:   `HMS Video Kit HDR Vivid`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HMS Video Kit HDR Vivid`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libNvEffectSdkCore.so`,
		Zh: SDKLangInfo{
			Label:   `美摄 SDK`,
			DevTeam: `美摄`,
		},
		En: SDKLangInfo{
			Label:   `Meishe SDK`,
			DevTeam: `美摄`,
		},
	},
	{
		Soname: `libbrlttywrap.so`,
		Zh: SDKLangInfo{
			Label:   `brltty`,
			DevTeam: `Debian`,
		},
		En: SDKLangInfo{
			Label:   `brltty`,
			DevTeam: `Debian`,
		},
	},
	{
		Soname: `libUE4.so`,
		Zh: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
		En: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
	},
	{
		Soname: `libmnnkitcore.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libjiagu_vip_x64.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libfile_lock_pg.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libphonon_fmod.so`,
		Zh: SDKLangInfo{
			Label:   `Steam Audio`,
			DevTeam: `Valve Corporation`,
		},
		En: SDKLangInfo{
			Label:   `Steam Audio`,
			DevTeam: `Valve Corporation`,
		},
	},
	{
		Soname: `libswresample.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libffmpegloader.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libzstandard_android.so`,
		Zh: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libxylinkAIEngine.so`,
		Zh: SDKLangInfo{
			Label:   `小鱼易联终端 SDK`,
			DevTeam: `小鱼易联`,
		},
		En: SDKLangInfo{
			Label:   `XYLink Terminal SDK`,
			DevTeam: `小鱼易联`,
		},
	},
	{
		Soname: `libagora-rtm-sdk-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libagora_ai_echo_cancellation_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libtbs.renderer_dex.jar.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libflutter_avif.so`,
		Zh: SDKLangInfo{
			Label:   `flutter_avif`,
			DevTeam: `yekeskin`,
		},
		En: SDKLangInfo{
			Label:   `flutter_avif`,
			DevTeam: `yekeskin`,
		},
	},
	{
		Soname: `libsgmiddletier.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsqliteJni.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack SQLite`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack SQLite`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbreakpad_cpp_helper.so`,
		Zh: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libjiagu_x86.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libNvStreamingSdkCore.so`,
		Zh: SDKLangInfo{
			Label:   `美摄 SDK`,
			DevTeam: `美摄`,
		},
		En: SDKLangInfo{
			Label:   `Meishe SDK`,
			DevTeam: `美摄`,
		},
	},
	{
		Soname: `libtobEmbedEncrypt.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libaterm.so`,
		Zh: SDKLangInfo{
			Label:   `aterminal`,
			DevTeam: `maoabc`,
		},
		En: SDKLangInfo{
			Label:   `aterminal`,
			DevTeam: `maoabc`,
		},
	},
	{
		Soname: `libtensorflowlite_gpu_jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libfreetypejni.so`,
		Zh: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
		En: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
	},
	{
		Soname: `libglide-webp.so`,
		Zh: SDKLangInfo{
			Label:   `GlideWebpDecoder`,
			DevTeam: `zjupure`,
		},
		En: SDKLangInfo{
			Label:   `GlideWebpDecoder`,
			DevTeam: `zjupure`,
		},
	},
	{
		Soname: `libentryexpro.so`,
		Zh: SDKLangInfo{
			Label:   `银联 SDK`,
			DevTeam: `银联`,
		},
		En: SDKLangInfo{
			Label:   `UnionPay SDK`,
			DevTeam: `银联`,
		},
	},
	{
		Soname: `libindoor.so`,
		Zh: SDKLangInfo{
			Label:   `百度 LBS`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu LBS`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libYTFaceDetector.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libsonic.so`,
		Zh: SDKLangInfo{
			Label:   `Sonic`,
			DevTeam: `waywardgeek`,
		},
		En: SDKLangInfo{
			Label:   `Sonic`,
			DevTeam: `waywardgeek`,
		},
	},
	{
		Soname: `libYTAGReflectLiveCheck.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libjpeg-turbo.so`,
		Zh: SDKLangInfo{
			Label:   `libjpeg-turbo`,
			DevTeam: `dcommander`,
		},
		En: SDKLangInfo{
			Label:   `libjpeg-turbo`,
			DevTeam: `dcommander`,
		},
	},
	{
		Soname: `libsecenh.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA Reinforcement`,
			DevTeam: `CFCA`,
		},
		En: SDKLangInfo{
			Label:   `CFCA 加固`,
			DevTeam: `CFCA`,
		},
	},
	{
		Soname: `libmegface_v5.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libmetasec_ml.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libedog.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `lib39285EFA.so`,
		Zh: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
		En: SDKLangInfo{
			Label:   `MSA SDK`,
			DevTeam: `移动安全联盟`,
		},
	},
	{
		Soname: `libsnpe-android.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libsscronet.so`,
		Zh: SDKLangInfo{
			Label:   `Cronet`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `Cronet`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `libturbomodulejsijni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libnmmvm.so`,
		Zh: SDKLangInfo{
			Label:   `nmmp`,
			DevTeam: `maoabc`,
		},
		En: SDKLangInfo{
			Label:   `nmmp`,
			DevTeam: `maoabc`,
		},
	},
	{
		Soname: `libnelprender.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libraw.so`,
		Zh: SDKLangInfo{
			Label:   `LibRaw`,
			DevTeam: `LibRaw LLC`,
		},
		En: SDKLangInfo{
			Label:   `LibRaw`,
			DevTeam: `LibRaw LLC`,
		},
	},
	{
		Soname: `libijiami_ijm.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libjeffmony.so`,
		Zh: SDKLangInfo{
			Label:   `PlayerSDK`,
			DevTeam: `JeffMony`,
		},
		En: SDKLangInfo{
			Label:   `PlayerSDK`,
			DevTeam: `JeffMony`,
		},
	},
	{
		Soname: `libsurface_util_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack Camera`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack Camera`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libcrashlytics-common.so`,
		Zh: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libjscinstance.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libmnn_jsi.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libandroidkeyboard.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libalicomphonenumberauthsdk-log-online-standard-release_alijtca_plus.so`,
		Zh: SDKLangInfo{
			Label:   `号码认证服务`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Number Authentication Service`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libkraken.so`,
		Zh: SDKLangInfo{
			Label:   `Kraken`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Kraken`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libframesequence.so`,
		Zh: SDKLangInfo{
			Label:   `FrameSequence`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `FrameSequence`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `librtbt828.so`,
		Zh: SDKLangInfo{
			Label:   `高德地图 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Amap SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libPerfgeniusApi.so`,
		Zh: SDKLangInfo{
			Label:   `PerfGenius`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `PerfGenius`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libmlkit_google_ocr_pipeline.so`,
		Zh: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ML Kit`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbrltty.so`,
		Zh: SDKLangInfo{
			Label:   `brltty`,
			DevTeam: `Debian`,
		},
		En: SDKLangInfo{
			Label:   `brltty`,
			DevTeam: `Debian`,
		},
	},
	{
		Soname: `libygsiyu.so`,
		Zh: SDKLangInfo{
			Label:   `iApp`,
			DevTeam: `iApp`,
		},
		En: SDKLangInfo{
			Label:   `iApp`,
			DevTeam: `iApp`,
		},
	},
	{
		Soname: `libmediastreamer_base.so`,
		Zh: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
		En: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
	},
	{
		Soname: `libreact_codegen_RNCViewPager.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libhiai_hcl_model_runtime.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libbaiduprotect.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libimage_processing_util_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack Camera`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack Camera`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbugsnag-plugin-android-anr.so`,
		Zh: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
		En: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
	},
	{
		Soname: `libmibrainsdk.so`,
		Zh: SDKLangInfo{
			Label:   `小爱 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiao Ai SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libswscale.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libopencv_java4.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `libsgavmp.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `librenderscript-toolkit.so`,
		Zh: SDKLangInfo{
			Label:   `RenderScript Intrinsics Replacement Toolkit`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `RenderScript Intrinsics Replacement Toolkit`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libimagepipeline_x86.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libtensorflowlite.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `librime_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Rime Core`,
			DevTeam: `Rime (create by lotem)`,
		},
		En: SDKLangInfo{
			Label:   `Rime Core`,
			DevTeam: `Rime (create by lotem)`,
		},
	},
	{
		Soname: `libweex_framework.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsgmainso.version.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `liboctopus_classifier.so`,
		Zh: SDKLangInfo{
			Label:   `liboctopus`,
			DevTeam: `opendesigndev`,
		},
		En: SDKLangInfo{
			Label:   `liboctopus`,
			DevTeam: `opendesigndev`,
		},
	},
	{
		Soname: `libmpaascpu.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libshadowhook.so`,
		Zh: SDKLangInfo{
			Label:   `ShadowHook`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `ShadowHook`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libsdkcore.so`,
		Zh: SDKLangInfo{
			Label:   `FinClip`,
			DevTeam: `凡泰极客`,
		},
		En: SDKLangInfo{
			Label:   `FinClip`,
			DevTeam: `凡泰极客`,
		},
	},
	{
		Soname: `libsig_native.so`,
		Zh: SDKLangInfo{
			Label:   `LibSig`,
			DevTeam: `Qix`,
		},
		En: SDKLangInfo{
			Label:   `LibSig`,
			DevTeam: `Qix`,
		},
	},
	{
		Soname: `libyaml.so`,
		Zh: SDKLangInfo{
			Label:   `libyaml`,
			DevTeam: `yaml`,
		},
		En: SDKLangInfo{
			Label:   `libyaml`,
			DevTeam: `yaml`,
		},
	},
	{
		Soname: `libmarsxlog.so`,
		Zh: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Mars`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libgvr.so`,
		Zh: SDKLangInfo{
			Label:   `Google VR SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google VR SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libtongdun.so`,
		Zh: SDKLangInfo{
			Label:   `同盾设备指纹`,
			DevTeam: `同盾科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Tongdun Device Fingerprint`,
			DevTeam: `同盾科技有限公司`,
		},
	},
	{
		Soname: `libreact_codegen_ReactSlider.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libMSDKCore.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libZegoExpressEngine.so`,
		Zh: SDKLangInfo{
			Label:   `即构极速视频 SDK`,
			DevTeam: `即构`,
		},
		En: SDKLangInfo{
			Label:   `Zego Express Video SDK`,
			DevTeam: `即构`,
		},
	},
	{
		Soname: `libjsijniprofiler.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `librhttp.so`,
		Zh: SDKLangInfo{
			Label:   `rhttp`,
			DevTeam: `tienisto.com`,
		},
		En: SDKLangInfo{
			Label:   `rhttp`,
			DevTeam: `tienisto.com`,
		},
	},
	{
		Soname: `libsmsdk.so`,
		Zh: SDKLangInfo{
			Label:   `数美 SDK`,
			DevTeam: `数美`,
		},
		En: SDKLangInfo{
			Label:   `Shumei SDK`,
			DevTeam: `数美`,
		},
	},
	{
		Soname: `libmnnpybridge.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libovpncli.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libffmpeg_resample.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `liblspatch.so`,
		Zh: SDKLangInfo{
			Label:   `LSPatch`,
			DevTeam: `LSPosed`,
		},
		En: SDKLangInfo{
			Label:   `LSPatch`,
			DevTeam: `LSPosed`,
		},
	},
	{
		Soname: `libhippybridge.so`,
		Zh: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Hippy`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libexpo-av.so`,
		Zh: SDKLangInfo{
			Label:   `Expo AV`,
			DevTeam: `Expo`,
		},
		En: SDKLangInfo{
			Label:   `Expo AV`,
			DevTeam: `Expo`,
		},
	},
	{
		Soname: `libsgnocaptcha.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `liblivestreaming.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `liblight-sdk.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯特效 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Special Effects SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `lib_burst_generated.so`,
		Zh: SDKLangInfo{
			Label:   `Burst`,
			DevTeam: `Unity`,
		},
		En: SDKLangInfo{
			Label:   `Burst`,
			DevTeam: `Unity`,
		},
	},
	{
		Soname: `libBugly-idasc.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libsqlite3_chaquopy.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libbangcle_crypto_tool.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libbuglybacktrace.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libtencentsm-lite.so`,
		Zh: SDKLangInfo{
			Label:   `Kona 国密套件`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Kona SM Suite`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libkoom-strip-dump.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libzip.so`,
		Zh: SDKLangInfo{
			Label:   `libzip`,
			DevTeam: `nih-at`,
		},
		En: SDKLangInfo{
			Label:   `libzip`,
			DevTeam: `nih-at`,
		},
	},
	{
		Soname: `libwg.so`,
		Zh: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
		En: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
	},
	{
		Soname: `libcom.baidu.zeus.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libzstd_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libproperty_get.so`,
		Zh: SDKLangInfo{
			Label:   `AntiFakerAndroidChecker`,
			DevTeam: `happylishang`,
		},
		En: SDKLangInfo{
			Label:   `AntiFakerAndroidChecker`,
			DevTeam: `happylishang`,
		},
	},
	{
		Soname: `libencrypt.so`,
		Zh: SDKLangInfo{
			Label:   `lib-encrypt`,
			DevTeam: `bihe0832`,
		},
		En: SDKLangInfo{
			Label:   `lib-encrypt`,
			DevTeam: `bihe0832`,
		},
	},
	{
		Soname: `libopuscodec.so`,
		Zh: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
		En: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
	},
	{
		Soname: `libsaturn.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云直播 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Live Streaming SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libjsiutil.so`,
		Zh: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
		En: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
	},
	{
		Soname: `libtbs.tbs_sdk_extension_dex.jar.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_codegen_Spec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libappsuit.so`,
		Zh: SDKLangInfo{
			Label:   `APKiD`,
			DevTeam: `rednaga`,
		},
		En: SDKLangInfo{
			Label:   `APKiD`,
			DevTeam: `rednaga`,
		},
	},
	{
		Soname: `libsqlitex.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libQPlayer.so`,
		Zh: SDKLangInfo{
			Label:   `PLDroidPlayer`,
			DevTeam: `七牛云`,
		},
		En: SDKLangInfo{
			Label:   `PLDroidPlayer`,
			DevTeam: `七牛云`,
		},
	},
	{
		Soname: `libllm_inference_engine_jni.so`,
		Zh: SDKLangInfo{
			Label:   `MediaPipe LLM Inference`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `MediaPipe LLM Inference`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libLive2DCubismCore.so`,
		Zh: SDKLangInfo{
			Label:   `Live2D Cubism SDK`,
			DevTeam: `Live2D`,
		},
		En: SDKLangInfo{
			Label:   `Live2D Cubism SDK`,
			DevTeam: `Live2D`,
		},
	},
	{
		Soname: `libMNN_CL.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libhiai_ir_build_aipp.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libmonodroid.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libBaiduSpeechAlgOTA.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libmodpdfium.so`,
		Zh: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `chaquopy.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libdart2cpp.so`,
		Zh: SDKLangInfo{
			Label:   `dart2cpp`,
			DevTeam: `Brendan Duncan`,
		},
		En: SDKLangInfo{
			Label:   `dart2cpp`,
			DevTeam: `Brendan Duncan`,
		},
	},
	{
		Soname: `librdk.so`,
		Zh: SDKLangInfo{
			Label:   `Foxit PDF SDK`,
			DevTeam: `Foxit`,
		},
		En: SDKLangInfo{
			Label:   `Foxit PDF SDK`,
			DevTeam: `Foxit`,
		},
	},
	{
		Soname: `libartc_engine.so`,
		Zh: SDKLangInfo{
			Label:   `HashLips Art Engine`,
			DevTeam: `chachohee`,
		},
		En: SDKLangInfo{
			Label:   `HashLips Art Engine`,
			DevTeam: `chachohee`,
		},
	},
	{
		Soname: `libwebp.so`,
		Zh: SDKLangInfo{
			Label:   `WebP 编解码器`,
			DevTeam: `WebM project`,
		},
		En: SDKLangInfo{
			Label:   `WebP Codec`,
			DevTeam: `WebM project`,
		},
	},
	{
		Soname: `libquickphrase.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `liblibrtlsdr.so`,
		Zh: SDKLangInfo{
			Label:   `librtlsdr`,
			DevTeam: `librtlsdr`,
		},
		En: SDKLangInfo{
			Label:   `librtlsdr`,
			DevTeam: `librtlsdr`,
		},
	},
	{
		Soname: `libcrypto_chaquopy.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libeidjni.so`,
		Zh: SDKLangInfo{
			Label:   `eID SDK`,
			DevTeam: `公安部第三研究所`,
		},
		En: SDKLangInfo{
			Label:   `eID SDK`,
			DevTeam: `公安部第三研究所`,
		},
	},
	{
		Soname: `libion.so`,
		Zh: SDKLangInfo{
			Label:   `ion`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ion`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnelpengine.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libPaddle.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libj2v8.so`,
		Zh: SDKLangInfo{
			Label:   `J2V8`,
			DevTeam: `Eclipse`,
		},
		En: SDKLangInfo{
			Label:   `J2V8`,
			DevTeam: `Eclipse`,
		},
	},
	{
		Soname: `libsentry_dumper.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libtmsqmp.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libfilterframework_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack Media`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack Media`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `librepublic.so`,
		Zh: SDKLangInfo{
			Label:   `RePublic`,
			DevTeam: `whulzz1993`,
		},
		En: SDKLangInfo{
			Label:   `RePublic`,
			DevTeam: `whulzz1993`,
		},
	},
	{
		Soname: `libtensorflowLite.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `librive_text.so`,
		Zh: SDKLangInfo{
			Label:   `Rive`,
			DevTeam: `Rive`,
		},
		En: SDKLangInfo{
			Label:   `Rive`,
			DevTeam: `Rive`,
		},
	},
	{
		Soname: `chaquopy_android.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libTUTKGlobalAPIs.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libSDL2.so`,
		Zh: SDKLangInfo{
			Label:   `libSDL`,
			DevTeam: `libsdl.org`,
		},
		En: SDKLangInfo{
			Label:   `libSDL`,
			DevTeam: `libsdl.org`,
		},
	},
	{
		Soname: `libfilament-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libalitanx-lib.so`,
		Zh: SDKLangInfo{
			Label:   `TanX`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `TanX`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libdobby.so`,
		Zh: SDKLangInfo{
			Label:   `Dobby`,
			DevTeam: `jmpews`,
		},
		En: SDKLangInfo{
			Label:   `Dobby`,
			DevTeam: `jmpews`,
		},
	},
	{
		Soname: `libRtsSDK.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libcolorpen.so`,
		Zh: SDKLangInfo{
			Label:   `即构极速视频 SDK`,
			DevTeam: `即构`,
		},
		En: SDKLangInfo{
			Label:   `Zego Express Video SDK`,
			DevTeam: `即构`,
		},
	},
	{
		Soname: `libopencc.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCC`,
			DevTeam: `BYVoid`,
		},
		En: SDKLangInfo{
			Label:   `OpenCC`,
			DevTeam: `BYVoid`,
		},
	},
	{
		Soname: `libexec_x86.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeVideoSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libImSDK.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云通信 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Communication SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libsqlite3.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libxplatform.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libMyGame.so`,
		Zh: SDKLangInfo{
			Label:   `cocos2d-x`,
			DevTeam: `cocos2d`,
		},
		En: SDKLangInfo{
			Label:   `cocos2d-x`,
			DevTeam: `cocos2d`,
		},
	},
	{
		Soname: `libarcsoft_lensstaindetection.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libmpaas_teclability.so`,
		Zh: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `mPaaS`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libcompat-limbo.so`,
		Zh: SDKLangInfo{
			Label:   `Limbo`,
			DevTeam: `limboemu`,
		},
		En: SDKLangInfo{
			Label:   `Limbo`,
			DevTeam: `limboemu`,
		},
	},
	{
		Soname: `libIMECore.so`,
		Zh: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libffmpegkit_abidetect.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpegKit`,
			DevTeam: `arthenica`,
		},
		En: SDKLangInfo{
			Label:   `FFmpegKit`,
			DevTeam: `arthenica`,
		},
	},
	{
		Soname: `lib_RongRTC_so.so`,
		Zh: SDKLangInfo{
			Label:   `融云 IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
		En: SDKLangInfo{
			Label:   `RongCloud IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
	},
	{
		Soname: `libhpplayae.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libbox.so`,
		Zh: SDKLangInfo{
			Label:   `sing-box`,
			DevTeam: `SagerNet`,
		},
		En: SDKLangInfo{
			Label:   `sing-box`,
			DevTeam: `SagerNet`,
		},
	},
	{
		Soname: `libhiai_binary_model_runtime.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libagora_dav1d_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libYTFaceAlignment.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `lbtoyger-release-alijtca_plus.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libRNQuickSQlite.so`,
		Zh: SDKLangInfo{
			Label:   `RNQuickSQlite`,
			DevTeam: `margelo`,
		},
		En: SDKLangInfo{
			Label:   `RNQuickSQlite`,
			DevTeam: `margelo`,
		},
	},
	{
		Soname: `libx5linker.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmibrainjni.so`,
		Zh: SDKLangInfo{
			Label:   `小爱 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiao Ai SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libantitrace.so`,
		Zh: SDKLangInfo{
			Label:   `EasyProtector`,
			DevTeam: `lamster2018`,
		},
		En: SDKLangInfo{
			Label:   `EasyProtector`,
			DevTeam: `lamster2018`,
		},
	},
	{
		Soname: `libconscrypt_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Conscrypt`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Conscrypt`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbreakpad_client.so`,
		Zh: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libreact_codegen_PagerView.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libkwai-android-base.so`,
		Zh: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
		En: SDKLangInfo{
			Label:   `KOOM`,
			DevTeam: `KwaiAppTeam`,
		},
	},
	{
		Soname: `libhev-socks5-tunnel.so`,
		Zh: SDKLangInfo{
			Label:   `HevSocks5Tunnel`,
			DevTeam: `heiher`,
		},
		En: SDKLangInfo{
			Label:   `HevSocks5Tunnel`,
			DevTeam: `heiher`,
		},
	},
	{
		Soname: `libwebpbackport.so`,
		Zh: SDKLangInfo{
			Label:   `WebP`,
			DevTeam: `Alexey Pelykh`,
		},
		En: SDKLangInfo{
			Label:   `WebP`,
			DevTeam: `Alexey Pelykh`,
		},
	},
	{
		Soname: `libreact_codegen_rnscreens.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libApkPatchLibrary.so`,
		Zh: SDKLangInfo{
			Label:   `SmartAppUpdates`,
			DevTeam: `cundong`,
		},
		En: SDKLangInfo{
			Label:   `SmartAppUpdates`,
			DevTeam: `cundong`,
		},
	},
	{
		Soname: `libswan-arch.so`,
		Zh: SDKLangInfo{
			Label:   `智能小程序`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Smart Mini Program`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libphonon.so`,
		Zh: SDKLangInfo{
			Label:   `Steam Audio`,
			DevTeam: `Valve Corporation`,
		},
		En: SDKLangInfo{
			Label:   `Steam Audio`,
			DevTeam: `Valve Corporation`,
		},
	},
	{
		Soname: `libtun2socks.so`,
		Zh: SDKLangInfo{
			Label:   `Tun2Socks`,
			DevTeam: `Jason Lyu（xjasonlyu）`,
		},
		En: SDKLangInfo{
			Label:   `Tun2Socks`,
			DevTeam: `Jason Lyu（xjasonlyu）`,
		},
	},
	{
		Soname: `libalivcffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `阿里云短视频 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Cloud Short Video SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libdav1d.so`,
		Zh: SDKLangInfo{
			Label:   `dav1d`,
			DevTeam: `videolan`,
		},
		En: SDKLangInfo{
			Label:   `dav1d`,
			DevTeam: `videolan`,
		},
	},
	{
		Soname: `libgraphics-core.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack Graphics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack Graphics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbrotli.so`,
		Zh: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libpunctuation.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libcri_mana_vpx.so`,
		Zh: SDKLangInfo{
			Label:   `CRI Sofdec2 Unity 插件`,
			DevTeam: `CRIWARE`,
		},
		En: SDKLangInfo{
			Label:   `CRI Sofdec2 Unity Plugin`,
			DevTeam: `CRIWARE`,
		},
	},
	{
		Soname: `libliteavsdk.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云短视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Short Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libumeng-spy.so`,
		Zh: SDKLangInfo{
			Label:   `移动统计分析`,
			DevTeam: `Umeng`,
		},
		En: SDKLangInfo{
			Label:   `Mobile Statistics Analysis`,
			DevTeam: `Umeng`,
		},
	},
	{
		Soname: `libpojavexec_awt.so`,
		Zh: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
		En: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
	},
	{
		Soname: `libreactperfloggerjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libblackbox.so`,
		Zh: SDKLangInfo{
			Label:   `BlackBox`,
			DevTeam: `FBlackBox`,
		},
		En: SDKLangInfo{
			Label:   `BlackBox`,
			DevTeam: `FBlackBox`,
		},
	},
	{
		Soname: `libapssdk.so`,
		Zh: SDKLangInfo{
			Label:   `高德 SDK`,
			DevTeam: `高德`,
		},
		En: SDKLangInfo{
			Label:   `Amap SDK`,
			DevTeam: `高德`,
		},
	},
	{
		Soname: `libtcc.so`,
		Zh: SDKLangInfo{
			Label:   `libtcc`,
			DevTeam: `Fabrice Bellard`,
		},
		En: SDKLangInfo{
			Label:   `libtcc`,
			DevTeam: `Fabrice Bellard`,
		},
	},
	{
		Soname: `librust.so`,
		Zh: SDKLangInfo{
			Label:   `Rust`,
			DevTeam: `Mozilla`,
		},
		En: SDKLangInfo{
			Label:   `Rust`,
			DevTeam: `Mozilla`,
		},
	},
	{
		Soname: `libtxffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯云短视频 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Cloud Short Video SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libYTCommonLiveness.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liboboe.so`,
		Zh: SDKLangInfo{
			Label:   `Oboe`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Oboe`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmatrix-opengl-leak.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libKF5ConfigGui_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libpuerts.so`,
		Zh: SDKLangInfo{
			Label:   `PuerTS`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PuerTS`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libjsinspector.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libfilament-texture-utils-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libsentry-record.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libonnxruntime4j_jni.so`,
		Zh: SDKLangInfo{
			Label:   `onnxruntime`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `onnxruntime`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `librlottie-image-loader.so`,
		Zh: SDKLangInfo{
			Label:   `rlottie`,
			DevTeam: `Samsung`,
		},
		En: SDKLangInfo{
			Label:   `rlottie`,
			DevTeam: `Samsung`,
		},
	},
	{
		Soname: `libbreakpad-core.so`,
		Zh: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Breakpad`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libcyaml.so`,
		Zh: SDKLangInfo{
			Label:   `libyaml`,
			DevTeam: `yaml`,
		},
		En: SDKLangInfo{
			Label:   `libyaml`,
			DevTeam: `yaml`,
		},
	},
	{
		Soname: `libtraceroute.so`,
		Zh: SDKLangInfo{
			Label:   `traceroute`,
			DevTeam: `ilyagrishkov`,
		},
		En: SDKLangInfo{
			Label:   `traceroute`,
			DevTeam: `ilyagrishkov`,
		},
	},
	{
		Soname: `libBifrost.so`,
		Zh: SDKLangInfo{
			Label:   `Bifrost`,
			DevTeam: `pmelsted`,
		},
		En: SDKLangInfo{
			Label:   `Bifrost`,
			DevTeam: `pmelsted`,
		},
	},
	{
		Soname: `libhyphenate_av.so`,
		Zh: SDKLangInfo{
			Label:   `环信 IM`,
			DevTeam: `环信`,
		},
		En: SDKLangInfo{
			Label:   `Easemob IM`,
			DevTeam: `环信`,
		},
	},
	{
		Soname: `libmegface_faceid.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libacra.so`,
		Zh: SDKLangInfo{
			Label:   `ACRA`,
			DevTeam: `ACRA`,
		},
		En: SDKLangInfo{
			Label:   `ACRA`,
			DevTeam: `ACRA`,
		},
	},
	{
		Soname: `libaot-mscorlib.dll.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libxguardian.so`,
		Zh: SDKLangInfo{
			Label:   `信鸽推送`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Xinge Push`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libagora_ci_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libprotobuf-lite.so`,
		Zh: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `protobuf`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libopvpnutil.so`,
		Zh: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
		En: SDKLangInfo{
			Label:   `OpenVPN`,
			DevTeam: `OpenVPN Inc.`,
		},
	},
	{
		Soname: `libttboringssl.so`,
		Zh: SDKLangInfo{
			Label:   `BoringSSL`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `BoringSSL`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libSecShell.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libaliyunaf.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libagora_segmentation_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libwlibandroid.so`,
		Zh: SDKLangInfo{
			Label:   `Wlib`,
			DevTeam: `Windr07`,
		},
		En: SDKLangInfo{
			Label:   `Wlib`,
			DevTeam: `Windr07`,
		},
	},
	{
		Soname: `libAliNNPython_zip.so`,
		Zh: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `AliNNPython`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libshadowhook_nothing.so`,
		Zh: SDKLangInfo{
			Label:   `ShadowHook`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `ShadowHook`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libxgVipSecurity.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯移动推送`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Mobile Push`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblogging.so`,
		Zh: SDKLangInfo{
			Label:   `logging`,
			DevTeam: `rsyslog`,
		},
		En: SDKLangInfo{
			Label:   `logging`,
			DevTeam: `rsyslog`,
		},
	},
	{
		Soname: `libmsx264.so`,
		Zh: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
		En: SDKLangInfo{
			Label:   `Mediastreamer2`,
			DevTeam: `Belledonne Communications`,
		},
	},
	{
		Soname: `libnative-tianmu-lib.so`,
		Zh: SDKLangInfo{
			Label:   `Tianmu SDK`,
			DevTeam: `杭州艾狄墨搏信息服务有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Tianmu SDK`,
			DevTeam: `杭州艾狄墨搏信息服务有限公司`,
		},
	},
	{
		Soname: `libzegoqueue.so`,
		Zh: SDKLangInfo{
			Label:   `即构极速视频 SDK`,
			DevTeam: `即构`,
		},
		En: SDKLangInfo{
			Label:   `Zego Express Video SDK`,
			DevTeam: `即构`,
		},
	},
	{
		Soname: `libNativeScript.so`,
		Zh: SDKLangInfo{
			Label:   `NativeScripts`,
			DevTeam: `OpenJS Foundation`,
		},
		En: SDKLangInfo{
			Label:   `NativeScripts`,
			DevTeam: `OpenJS Foundation`,
		},
	},
	{
		Soname: `libPixUI_Unity.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtess.so`,
		Zh: SDKLangInfo{
			Label:   `Tesseract`,
			DevTeam: `Tesseract`,
		},
		En: SDKLangInfo{
			Label:   `Tesseract`,
			DevTeam: `Tesseract`,
		},
	},
	{
		Soname: `libmemchunk.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libapmioFake.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯客户端性能分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Client Performance Analysis`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libhiai_c_def.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libutility.so`,
		Zh: SDKLangInfo{
			Label:   `微博 SDK`,
			DevTeam: `Weibo`,
		},
		En: SDKLangInfo{
			Label:   `Weibo SDK`,
			DevTeam: `Weibo`,
		},
	},
	{
		Soname: `librnskia.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Skia`,
			DevTeam: `Shopify`,
		},
		En: SDKLangInfo{
			Label:   `React Native Skia`,
			DevTeam: `Shopify`,
		},
	},
	{
		Soname: `libaria2c.so`,
		Zh: SDKLangInfo{
			Label:   `libaria2`,
			DevTeam: `aria2`,
		},
		En: SDKLangInfo{
			Label:   `libaria2`,
			DevTeam: `aria2`,
		},
	},
	{
		Soname: `libNativeImageProcessor.so`,
		Zh: SDKLangInfo{
			Label:   `AndroidPhotoFilters`,
			DevTeam: `Zomato`,
		},
		En: SDKLangInfo{
			Label:   `AndroidPhotoFilters`,
			DevTeam: `Zomato`,
		},
	},
	{
		Soname: `libpymodules.so`,
		Zh: SDKLangInfo{
			Label:   `python-for-android`,
			DevTeam: `kivy`,
		},
		En: SDKLangInfo{
			Label:   `python-for-android`,
			DevTeam: `kivy`,
		},
	},
	{
		Soname: `libbd_easr_s1_merge_normal_20151216.dat.so`,
		Zh: SDKLangInfo{
			Label:   `语音识别 SDK`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Speech Recognition SDK`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libztcodec2.so`,
		Zh: SDKLangInfo{
			Label:   `阿里云语音服务 SDK（NLSClinet）`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Cloud Voice Service SDK (NLSClinet)`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libpcs_tensorflow_jni.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libtrace-canary.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbugsnag-ndk.so`,
		Zh: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
		En: SDKLangInfo{
			Label:   `Bugsnag`,
			DevTeam: `Bugsnag`,
		},
	},
	{
		Soname: `libnative-imagetranscoder.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libBARDumix.so`,
		Zh: SDKLangInfo{
			Label:   `BARDUmix`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `BARDUmix`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libclash.so`,
		Zh: SDKLangInfo{
			Label:   `Clash Core`,
			DevTeam: `Dreamacro`,
		},
		En: SDKLangInfo{
			Label:   `Clash Core`,
			DevTeam: `Dreamacro`,
		},
	},
	{
		Soname: `libcrashlytics.so`,
		Zh: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Crashlytics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `liblogger.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libcipherdb.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libreactnativeutilsjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native App Utils`,
			DevTeam: `Stumble App`,
		},
		En: SDKLangInfo{
			Label:   `React Native App Utils`,
			DevTeam: `Stumble App`,
		},
	},
	{
		Soname: `libImgRecognition.so`,
		Zh: SDKLangInfo{
			Label:   `DuMix`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `DuMix`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libmonosgen-2.0.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libtxmapengine.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbutuqrcode.so`,
		Zh: SDKLangInfo{
			Label:   `Butu QR SDK`,
			DevTeam: `钚兔团队`,
		},
		En: SDKLangInfo{
			Label:   `钚兔扫一扫 SDK`,
			DevTeam: `钚兔团队`,
		},
	},
	{
		Soname: `libbuffer_pg.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libchrome.so`,
		Zh: SDKLangInfo{
			Label:   `libchrome`,
			DevTeam: `Chromium`,
		},
		En: SDKLangInfo{
			Label:   `libchrome`,
			DevTeam: `Chromium`,
		},
	},
	{
		Soname: `libWAJ.so`,
		Zh: SDKLangInfo{
			Label:   `WAJ`,
			DevTeam: `Rust`,
		},
		En: SDKLangInfo{
			Label:   `WAJ`,
			DevTeam: `Rust`,
		},
	},
	{
		Soname: `libboost_multidex.so`,
		Zh: SDKLangInfo{
			Label:   `BoostMultiDex`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BoostMultiDex`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libtensorflow_inference.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libmace.so`,
		Zh: SDKLangInfo{
			Label:   `MACE`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `MACE`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `liblinphone.so`,
		Zh: SDKLangInfo{
			Label:   `Linphone`,
			DevTeam: `Linphone`,
		},
		En: SDKLangInfo{
			Label:   `Linphone`,
			DevTeam: `Linphone`,
		},
	},
	{
		Soname: `libadcolony.so`,
		Zh: SDKLangInfo{
			Label:   `AdColony`,
			DevTeam: `AdColony`,
		},
		En: SDKLangInfo{
			Label:   `AdColony`,
			DevTeam: `AdColony`,
		},
	},
	{
		Soname: `libpnc-crypto.so`,
		Zh: SDKLangInfo{
			Label:   `梆梆安全`,
			DevTeam: `梆梆安全`,
		},
		En: SDKLangInfo{
			Label:   `Bangbang Security`,
			DevTeam: `梆梆安全`,
		},
	},
	{
		Soname: `libluajava.so`,
		Zh: SDKLangInfo{
			Label:   `AndroLua`,
			DevTeam: `mkottman`,
		},
		En: SDKLangInfo{
			Label:   `AndroLua`,
			DevTeam: `mkottman`,
		},
	},
	{
		Soname: `libagora-ffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libagora-crypto.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libhuawei_arengine_ndk.so`,
		Zh: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `AR Engine`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libkernelu4_7z_uc.so`,
		Zh: SDKLangInfo{
			Label:   `U4 内核`,
			DevTeam: `UC`,
		},
		En: SDKLangInfo{
			Label:   `U4 内核`,
			DevTeam: `UC`,
		},
	},
	{
		Soname: `libsandhook64.so`,
		Zh: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
		En: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
	},
	{
		Soname: `libopenssl_init.so`,
		Zh: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
		En: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
	},
	{
		Soname: `libGXAnalyzeCore.so`,
		Zh: SDKLangInfo{
			Label:   `GaiaX`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `GaiaX`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libjiagu_x64.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libswanKV.so`,
		Zh: SDKLangInfo{
			Label:   `智能小程序`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Smart Mini Program`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libvideo_dec.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libgcanvas_runtime_jsi.so`,
		Zh: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libswappyVk.so`,
		Zh: SDKLangInfo{
			Label:   `Android 游戏开发工具包帧率调整`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Android Game Development Kit Frame Pacing`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbmob.so`,
		Zh: SDKLangInfo{
			Label:   `Bmob 后端云`,
			DevTeam: `Bmob`,
		},
		En: SDKLangInfo{
			Label:   `Bmob Backend Cloud`,
			DevTeam: `Bmob`,
		},
	},
	{
		Soname: `libsgmain.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libspell.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libxnn_core.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libagora_ai_denoise_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libreactnativeblob.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libsketchology_native.so`,
		Zh: SDKLangInfo{
			Label:   `Google Ink (Sketchology)`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google Ink (Sketchology)`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libhermes.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libffmpegJNI.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libxp2p.so`,
		Zh: SDKLangInfo{
			Label:   `libp2p`,
			DevTeam: `Protocol Labs`,
		},
		En: SDKLangInfo{
			Label:   `libp2p`,
			DevTeam: `Protocol Labs`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.silk_codec.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libdartcv.so`,
		Zh: SDKLangInfo{
			Label:   `DartCV`,
			DevTeam: `rainyl`,
		},
		En: SDKLangInfo{
			Label:   `DartCV`,
			DevTeam: `rainyl`,
		},
	},
	{
		Soname: `libYTFaceTrackPro2.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtensorflowlite_c.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libapminsightb.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libMNN_Legacy.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libjsengine-api.so`,
		Zh: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
		En: SDKLangInfo{
			Label:   `LibJS`,
			DevTeam: `linusg`,
		},
	},
	{
		Soname: `libquicklogin.so`,
		Zh: SDKLangInfo{
			Label:   `极光认证 SDK`,
			DevTeam: `极光`,
		},
		En: SDKLangInfo{
			Label:   `Aurora Authentication SDK`,
			DevTeam: `极光`,
		},
	},
	{
		Soname: `libreactnative.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libgesturehandler.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-gesture-handler`,
			DevTeam: `software-mansion`,
		},
		En: SDKLangInfo{
			Label:   `react-native-gesture-handler`,
			DevTeam: `software-mansion`,
		},
	},
	{
		Soname: `libgodot_android.so`,
		Zh: SDKLangInfo{
			Label:   `Godot Engine`,
			DevTeam: `Godot Developers`,
		},
		En: SDKLangInfo{
			Label:   `Godot Engine`,
			DevTeam: `Godot Developers`,
		},
	},
	{
		Soname: `libffavc.so`,
		Zh: SDKLangInfo{
			Label:   `ffavc`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `ffavc`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libhiai_enhance.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libMtaNativeCrash.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯移动分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Mobile Analytics`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libheif.so`,
		Zh: SDKLangInfo{
			Label:   `libheif`,
			DevTeam: `struktur AG`,
		},
		En: SDKLangInfo{
			Label:   `libheif`,
			DevTeam: `struktur AG`,
		},
	},
	{
		Soname: `libreact_codegen_BVLinearGradientSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libyuv-decoder.so`,
		Zh: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmono.so`,
		Zh: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `libweexjssr.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libiconv.so`,
		Zh: SDKLangInfo{
			Label:   `iconv`,
			DevTeam: `GNU`,
		},
		En: SDKLangInfo{
			Label:   `iconv`,
			DevTeam: `GNU`,
		},
	},
	{
		Soname: `libP2PTunnelAPIs.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libxloader.so`,
		Zh: SDKLangInfo{
			Label:   `Naga Reinforcement`,
			DevTeam: `娜迦科技`,
		},
		En: SDKLangInfo{
			Label:   `娜迦加固`,
			DevTeam: `娜迦科技`,
		},
	},
	{
		Soname: `libsandhook.so`,
		Zh: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
		En: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
	},
	{
		Soname: `libmatrix_resource_canary.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_featureflags.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libMSDK.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmatrix-jnihook.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libRongIMLib.so`,
		Zh: SDKLangInfo{
			Label:   `融云 IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
		En: SDKLangInfo{
			Label:   `RongCloud IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
	},
	{
		Soname: `libSnpeHtpPrepare.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libcrashsdk.so`,
		Zh: SDKLangInfo{
			Label:   `岳鹰全景监控`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Yueying Panoramic Monitoring`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libopencv.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `libtbsqimei_v8a.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔终端 ID 精准识别体系`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `QIMEI SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libPixFFmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libuimanagerjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libopus.so`,
		Zh: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
		En: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
	},
	{
		Soname: `libworklets.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Worklets`,
			DevTeam: `Software Mansion`,
		},
		En: SDKLangInfo{
			Label:   `React Native Worklets`,
			DevTeam: `Software Mansion`,
		},
	},
	{
		Soname: `libagora_screen_capture_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libbeaconid.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Lighthouse SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libagora_full_audio_format_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `librealm-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
		En: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
	},
	{
		Soname: `libfdk-aac.so`,
		Zh: SDKLangInfo{
			Label:   `opencore-amr`,
			DevTeam: `mstorsjo`,
		},
		En: SDKLangInfo{
			Label:   `opencore-amr`,
			DevTeam: `mstorsjo`,
		},
	},
	{
		Soname: `libapminsighta.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libtuanjie.so`,
		Zh: SDKLangInfo{
			Label:   `团结引擎`,
			DevTeam: `Unity 中国`,
		},
		En: SDKLangInfo{
			Label:   `Tuanjie Engine`,
			DevTeam: `Unity China`,
		},
	},
	{
		Soname: `libgcanvas.so`,
		Zh: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libjnidispatch.so`,
		Zh: SDKLangInfo{
			Label:   `Java Native Access`,
			DevTeam: `Java Native Access`,
		},
		En: SDKLangInfo{
			Label:   `Java Native Access`,
			DevTeam: `Java Native Access`,
		},
	},
	{
		Soname: `libswappy.so`,
		Zh: SDKLangInfo{
			Label:   `Android 游戏开发工具包帧率调整`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Android Game Development Kit Frame Pacing`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libIOTCAPIs.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libpolyvsdl.so`,
		Zh: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
		En: SDKLangInfo{
			Label:   `Polyv HTML5 Player`,
			DevTeam: `Polyv`,
		},
	},
	{
		Soname: `libbugly_dumper.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libcri_ware_unity.so`,
		Zh: SDKLangInfo{
			Label:   `CRI ADX2 Unity 插件`,
			DevTeam: `CRIWARE`,
		},
		En: SDKLangInfo{
			Label:   `CRI ADX2 Unity Plugin`,
			DevTeam: `CRIWARE`,
		},
	},
	{
		Soname: `libarcsoft_pic_trace_video.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libwg-go.so`,
		Zh: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
		En: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
	},
	{
		Soname: `libMMANDKSignature.so`,
		Zh: SDKLangInfo{
			Label:   `中国无线营销联盟通用 SDK`,
			DevTeam: `中国无线营销联盟`,
		},
		En: SDKLangInfo{
			Label:   `China Wireless Marketing Alliance General SDK`,
			DevTeam: `中国无线营销联盟`,
		},
	},
	{
		Soname: `libwcdb.so`,
		Zh: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libzwrapper.so`,
		Zh: SDKLangInfo{
			Label:   `ZWrapper`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `ZWrapper`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbspatch.so`,
		Zh: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `bspatch`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libreactjsexecutor.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `liblsplant.so`,
		Zh: SDKLangInfo{
			Label:   `LSPlant`,
			DevTeam: `LSPosed`,
		},
		En: SDKLangInfo{
			Label:   `LSPlant`,
			DevTeam: `LSPosed`,
		},
	},
	{
		Soname: `libmapbox-gl.so`,
		Zh: SDKLangInfo{
			Label:   `Mapbox GL Native`,
			DevTeam: `Mapbox`,
		},
		En: SDKLangInfo{
			Label:   `Mapbox GL Native`,
			DevTeam: `Mapbox`,
		},
	},
	{
		Soname: `libmiot_patch.so`,
		Zh: SDKLangInfo{
			Label:   `小米 IoT SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiaomi IoT SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libsqlite3x.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libBugly_Native_idasc.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libandroidlame.so`,
		Zh: SDKLangInfo{
			Label:   `AndroidLame`,
			DevTeam: `naman14`,
		},
		En: SDKLangInfo{
			Label:   `AndroidLame`,
			DevTeam: `naman14`,
		},
	},
	{
		Soname: `libEncryptorP.so`,
		Zh: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `Pangle SDK`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libMNN.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libwg-quick.so`,
		Zh: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
		En: SDKLangInfo{
			Label:   `WireGuard`,
			DevTeam: `Jason A. Donenfeld`,
		},
	},
	{
		Soname: `libhfyuv.so`,
		Zh: SDKLangInfo{
			Label:   `libhfyuv`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libhfyuv`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libjiagu_a64.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libhermes-inspector.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libpngquantandroid.so`,
		Zh: SDKLangInfo{
			Label:   `pngquant-android`,
			DevTeam: `Nic Dahlquist`,
		},
		En: SDKLangInfo{
			Label:   `pngquant-android`,
			DevTeam: `Nic Dahlquist`,
		},
	},
	{
		Soname: `libhigh-available.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libmediautil_v8.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `libreact_utilities.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libtbs.MANIFEST.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtrojan-go.so`,
		Zh: SDKLangInfo{
			Label:   `Trojan-Go`,
			DevTeam: `p4gefau1t`,
		},
		En: SDKLangInfo{
			Label:   `Trojan-Go`,
			DevTeam: `p4gefau1t`,
		},
	},
	{
		Soname: `libconceal.so`,
		Zh: SDKLangInfo{
			Label:   `Conceal`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Conceal`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libil2cpp.so`,
		Zh: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `Unity`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `libagora-chat-sdk.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libmatrix_hprof_analyzer.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libgtcore.so`,
		Zh: SDKLangInfo{
			Label:   `极光认证 SDK`,
			DevTeam: `极光`,
		},
		En: SDKLangInfo{
			Label:   `Aurora Authentication SDK`,
			DevTeam: `极光`,
		},
	},
	{
		Soname: `libreactnativejni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libio-canary.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libun7zip.so`,
		Zh: SDKLangInfo{
			Label:   `Un7zip`,
			DevTeam: `hzy3774`,
		},
		En: SDKLangInfo{
			Label:   `Un7zip`,
			DevTeam: `hzy3774`,
		},
	},
	{
		Soname: `libMNN_NPU.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libffmpegkit.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpegKit`,
			DevTeam: `arthenica`,
		},
		En: SDKLangInfo{
			Label:   `FFmpegKit`,
			DevTeam: `arthenica`,
		},
	},
	{
		Soname: `libtpnsSecurity.so`,
		Zh: SDKLangInfo{
			Label:   `信鸽推送`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Xinge Push`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libobjectbox-jni.so`,
		Zh: SDKLangInfo{
			Label:   `ObjectBox`,
			DevTeam: `ObjectBox`,
		},
		En: SDKLangInfo{
			Label:   `ObjectBox`,
			DevTeam: `ObjectBox`,
		},
	},
	{
		Soname: `libopenssl.so`,
		Zh: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
		En: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
	},
	{
		Soname: `libWTF.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libweexjsb.so`,
		Zh: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Weex`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libzkfv_tj.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libzbar.so`,
		Zh: SDKLangInfo{
			Label:   `ZBar`,
			DevTeam: `spadix`,
		},
		En: SDKLangInfo{
			Label:   `ZBar`,
			DevTeam: `spadix`,
		},
	},
	{
		Soname: `libspeexdsp.so`,
		Zh: SDKLangInfo{
			Label:   `Speex`,
			DevTeam: `Xiph.Org Foundation`,
		},
		En: SDKLangInfo{
			Label:   `Speex`,
			DevTeam: `Xiph.Org Foundation`,
		},
	},
	{
		Soname: `libwz265.so`,
		Zh: SDKLangInfo{
			Label:   `Intelligent Video Encoding Engine`,
			DevTeam: `微帧/Visionular`,
		},
		En: SDKLangInfo{
			Label:   `视频智能编码引擎`,
			DevTeam: `微帧/Visionular`,
		},
	},
	{
		Soname: `libMSMNN.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libagora_video_decoder_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `librealmc.so`,
		Zh: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
		En: SDKLangInfo{
			Label:   `Realm`,
			DevTeam: `Realm`,
		},
	},
	{
		Soname: `libandroidnotification.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 for Android`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libPdfium.so`,
		Zh: SDKLangInfo{
			Label:   `PDFium`,
			DevTeam: `PDFium`,
		},
		En: SDKLangInfo{
			Label:   `PDFium`,
			DevTeam: `PDFium`,
		},
	},
	{
		Soname: `libgcanvas_runtime.so`,
		Zh: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `GCanvas`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libAPMSignalOOM.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libfolly_futures.so`,
		Zh: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Folly`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libdexpatch.so`,
		Zh: SDKLangInfo{
			Label:   `DexPatch`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `DexPatch`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libxhook.so`,
		Zh: SDKLangInfo{
			Label:   `xHook`,
			DevTeam: `iQIYI`,
		},
		En: SDKLangInfo{
			Label:   `xHook`,
			DevTeam: `iQIYI`,
		},
	},
	{
		Soname: `libreact_cxxreactpackage.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libopencv_world.so`,
		Zh: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
		En: SDKLangInfo{
			Label:   `OpenCV`,
			DevTeam: `OpenCV`,
		},
	},
	{
		Soname: `libleptonica.so`,
		Zh: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
		En: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
	},
	{
		Soname: `libspeex.so`,
		Zh: SDKLangInfo{
			Label:   `Speex`,
			DevTeam: `Xiph.Org Foundation`,
		},
		En: SDKLangInfo{
			Label:   `Speex`,
			DevTeam: `Xiph.Org Foundation`,
		},
	},
	{
		Soname: `libvlcjni.so`,
		Zh: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
		En: SDKLangInfo{
			Label:   `LibVLC`,
			DevTeam: `VideoLAN`,
		},
	},
	{
		Soname: `libvorbis.so`,
		Zh: SDKLangInfo{
			Label:   `Vorbis`,
			DevTeam: `Xiph.Org`,
		},
		En: SDKLangInfo{
			Label:   `Vorbis`,
			DevTeam: `Xiph.Org`,
		},
	},
	{
		Soname: `libbrotlienc.so`,
		Zh: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnode.so`,
		Zh: SDKLangInfo{
			Label:   `Node.js`,
			DevTeam: `plenluno`,
		},
		En: SDKLangInfo{
			Label:   `Node.js`,
			DevTeam: `plenluno`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeSvgSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libbarhopper_qr_only_jni.so`,
		Zh: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libhardcoder.so`,
		Zh: SDKLangInfo{
			Label:   `HardCoder`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `HardCoder`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `liblofelt_sdk.so`,
		Zh: SDKLangInfo{
			Label:   `Lofelt SDK`,
			DevTeam: `Lofelt`,
		},
		En: SDKLangInfo{
			Label:   `Lofelt SDK`,
			DevTeam: `Lofelt`,
		},
	},
	{
		Soname: `libgpuimage-library.so`,
		Zh: SDKLangInfo{
			Label:   `GPUImage`,
			DevTeam: `CyberAgent`,
		},
		En: SDKLangInfo{
			Label:   `GPUImage`,
			DevTeam: `CyberAgent`,
		},
	},
	{
		Soname: `libavfilter.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libxnn.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libdatastore_shared_counter.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack DataStore`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack DataStore`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libbyteplusrtc.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `librclone.so`,
		Zh: SDKLangInfo{
			Label:   `Rclone`,
			DevTeam: `Rclone`,
		},
		En: SDKLangInfo{
			Label:   `Rclone`,
			DevTeam: `Rclone`,
		},
	},
	{
		Soname: `libYTCommon.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libTencentSM.so`,
		Zh: SDKLangInfo{
			Label:   `国密套件`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `SM Suite`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libhpatchz.so`,
		Zh: SDKLangInfo{
			Label:   `HDiffPatch`,
			DevTeam: `sisong`,
		},
		En: SDKLangInfo{
			Label:   `HDiffPatch`,
			DevTeam: `sisong`,
		},
	},
	{
		Soname: `libaria2.so`,
		Zh: SDKLangInfo{
			Label:   `aria2`,
			DevTeam: `aria2`,
		},
		En: SDKLangInfo{
			Label:   `aria2`,
			DevTeam: `aria2`,
		},
	},
	{
		Soname: `libsqlite.so`,
		Zh: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
		En: SDKLangInfo{
			Label:   `SQLite`,
			DevTeam: `SQLite`,
		},
	},
	{
		Soname: `libreact_codegen_rnasyncstorage.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libijmDataEncryption_x86.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libvcap_core.so`,
		Zh: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
		En: SDKLangInfo{
			Label:   `VCAP`,
			DevTeam: `vivo`,
		},
	},
	{
		Soname: `libquickjs_abtest.so`,
		Zh: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
		En: SDKLangInfo{
			Label:   `QuickJS`,
			DevTeam: `Fabrice Bellard & Charlie Gordon`,
		},
	},
	{
		Soname: `libai_fmk_dnnacl.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libXNet.so`,
		Zh: SDKLangInfo{
			Label:   `XNet`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `XNet`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libreact_featureflagsjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libSNPE.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libSnpeCpu.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libfin-yuvutil.so`,
		Zh: SDKLangInfo{
			Label:   `FinClip`,
			DevTeam: `凡泰极客`,
		},
		En: SDKLangInfo{
			Label:   `FinClip`,
			DevTeam: `凡泰极客`,
		},
	},
	{
		Soname: `libreact_utils.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libyuv.so`,
		Zh: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libmegface.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `libfullwidth.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libmatrix-memguard.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libgltfio-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Filament`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `librive-android.so`,
		Zh: SDKLangInfo{
			Label:   `Rive`,
			DevTeam: `Rive`,
		},
		En: SDKLangInfo{
			Label:   `Rive`,
			DevTeam: `Rive`,
		},
	},
	{
		Soname: `libitophttpdns.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯游戏 MSDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Game MSDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libMNN_Express.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libvpx.so`,
		Zh: SDKLangInfo{
			Label:   `WebM VP8/VP9`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `WebM VP8/VP9`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libFcitx5Utils.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libreact_nativemodule_core.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Reanimated`,
			DevTeam: `software-mansion`,
		},
		En: SDKLangInfo{
			Label:   `React Native Reanimated`,
			DevTeam: `software-mansion`,
		},
	},
	{
		Soname: `libxphttpclientex.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmatrix-stack-tracer.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libfabricjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libKF5GuiAddons_arm64-v8a.so`,
		Zh: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
		En: SDKLangInfo{
			Label:   `KDE Framework`,
			DevTeam: `KDE`,
		},
	},
	{
		Soname: `libBugly_Native.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libtbs.libtbsqimei.so.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔终端 ID 精准识别体系`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `QIMEI SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libpojavexec.so`,
		Zh: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
		En: SDKLangInfo{
			Label:   `PojavLauncher`,
			DevTeam: `PojavLauncherTeam`,
		},
	},
	{
		Soname: `libmaesdk.so`,
		Zh: SDKLangInfo{
			Label:   `1DS`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `1DS`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libkrisp_noise_cancellation.so`,
		Zh: SDKLangInfo{
			Label:   `Krisp AI 语音 SDK`,
			DevTeam: `krisp`,
		},
		En: SDKLangInfo{
			Label:   `Krisp AI Voice SDK`,
			DevTeam: `krisp`,
		},
	},
	{
		Soname: `libASPEngineN.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libreact_codegen_ReactNativeWebviewSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libjlottie.so`,
		Zh: SDKLangInfo{
			Label:   `jLottie`,
			DevTeam: `LottieFiles`,
		},
		En: SDKLangInfo{
			Label:   `jLottie`,
			DevTeam: `LottieFiles`,
		},
	},
	{
		Soname: `libtesseract.so`,
		Zh: SDKLangInfo{
			Label:   `Tesseract`,
			DevTeam: `Tesseract`,
		},
		En: SDKLangInfo{
			Label:   `Tesseract`,
			DevTeam: `Tesseract`,
		},
	},
	{
		Soname: `libthreadfix.so`,
		Zh: SDKLangInfo{
			Label:   `Threadfix`,
			DevTeam: `Coalfire`,
		},
		En: SDKLangInfo{
			Label:   `Threadfix`,
			DevTeam: `Coalfire`,
		},
	},
	{
		Soname: `libsnpe_dsp_skel.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libblur.so`,
		Zh: SDKLangInfo{
			Label:   `StackBlur`,
			DevTeam: `Enrique López Mañas（kikoso）`,
		},
		En: SDKLangInfo{
			Label:   `StackBlur`,
			DevTeam: `Enrique López Mañas（kikoso）`,
		},
	},
	{
		Soname: `libfreetype.so`,
		Zh: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
		En: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
	},
	{
		Soname: `libvideo_enc.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libgmm-jni.so`,
		Zh: SDKLangInfo{
			Label:   `Google Map SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google Map SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnertc_sdk.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libsgmisc.so`,
		Zh: SDKLangInfo{
			Label:   `阿里聚安全`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Alibaba Ju Security`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libtdjni.so`,
		Zh: SDKLangInfo{
			Label:   `TDLib`,
			DevTeam: `Telegram`,
		},
		En: SDKLangInfo{
			Label:   `TDLib`,
			DevTeam: `Telegram`,
		},
	},
	{
		Soname: `libhermes-executor-common-debug.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libijmDataEncryption_arm64.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libancbase_rt_bokeh.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains_skel.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libhiai_ir_build.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libtquic.so`,
		Zh: SDKLangInfo{
			Label:   `TQUIC`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `TQUIC`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtensorflowlite_jni_gms_client.so`,
		Zh: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
		En: SDKLangInfo{
			Label:   `LiteRT`,
			DevTeam: `TensorFlow`,
		},
	},
	{
		Soname: `libxa-internal-api.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libjuicebox_sdk_jni.so`,
		Zh: SDKLangInfo{
			Label:   `Juicebox SDK`,
			DevTeam: `Juicebox Systems`,
		},
		En: SDKLangInfo{
			Label:   `Juicebox SDK`,
			DevTeam: `Juicebox Systems`,
		},
	},
	{
		Soname: `libmonobdwgc-2.0.so`,
		Zh: SDKLangInfo{
			Label:   `Unity Mono`,
			DevTeam: `Unity`,
		},
		En: SDKLangInfo{
			Label:   `Unity Mono`,
			DevTeam: `Unity`,
		},
	},
	{
		Soname: `libtbsqimei.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔终端 ID 精准识别体系`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `QIMEI SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libzstd.so`,
		Zh: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Zstandard`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libsobot.so`,
		Zh: SDKLangInfo{
			Label:   `智齿客服 SDK`,
			DevTeam: `智齿科技`,
		},
		En: SDKLangInfo{
			Label:   `Zhichi Customer Service SDK`,
			DevTeam: `智齿科技`,
		},
	},
	{
		Soname: `libksyapm.so`,
		Zh: SDKLangInfo{
			Label:   `金山云短视频编辑 SDK`,
			DevTeam: `Kingsoft`,
		},
		En: SDKLangInfo{
			Label:   `Kingsoft Cloud Short Video Editing SDK`,
			DevTeam: `Kingsoft`,
		},
	},
	{
		Soname: `liblog4cpp.so`,
		Zh: SDKLangInfo{
			Label:   `Log for C++`,
			DevTeam: `bastiaan, sanchouss_, waffel`,
		},
		En: SDKLangInfo{
			Label:   `Log for C++`,
			DevTeam: `bastiaan, sanchouss_, waffel`,
		},
	},
	{
		Soname: `libfuse.so`,
		Zh: SDKLangInfo{
			Label:   `libfuse`,
			DevTeam: `libfuse`,
		},
		En: SDKLangInfo{
			Label:   `libfuse`,
			DevTeam: `libfuse`,
		},
	},
	{
		Soname: `libagora_fd_extension.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libmnncore.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libxylinkAIEngineImpOnVulkan.so`,
		Zh: SDKLangInfo{
			Label:   `小鱼易联终端 SDK`,
			DevTeam: `小鱼易联`,
		},
		En: SDKLangInfo{
			Label:   `XYLink Terminal SDK`,
			DevTeam: `小鱼易联`,
		},
	},
	{
		Soname: `libgoogle-ocrclient-v3.so`,
		Zh: SDKLangInfo{
			Label:   `Google’s Drive API v3`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google’s Drive API v3`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libagora_fdkaac.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libgl4es.so`,
		Zh: SDKLangInfo{
			Label:   `GL4ES`,
			DevTeam: `ptitSeb`,
		},
		En: SDKLangInfo{
			Label:   `GL4ES`,
			DevTeam: `ptitSeb`,
		},
	},
	{
		Soname: `libapms_ndk_anr.so`,
		Zh: SDKLangInfo{
			Label:   `AppGallery Connect APM`,
			DevTeam: `Huawei`,
		},
		En: SDKLangInfo{
			Label:   `AppGallery Connect APM`,
			DevTeam: `Huawei`,
		},
	},
	{
		Soname: `libtbs.pb_dex.jar.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmegface.ocr.identify.so`,
		Zh: SDKLangInfo{
			Label:   `旷视 SDK`,
			DevTeam: `旷视`,
		},
		En: SDKLangInfo{
			Label:   `Megvii SDK`,
			DevTeam: `旷视`,
		},
	},
	{
		Soname: `librongcloud_xcrash_dumper.so`,
		Zh: SDKLangInfo{
			Label:   `融云 IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
		En: SDKLangInfo{
			Label:   `RongCloud IM SDK`,
			DevTeam: `融云 RongCloud`,
		},
	},
	{
		Soname: `libcocos2djs.so`,
		Zh: SDKLangInfo{
			Label:   `cocos2d-x`,
			DevTeam: `cocos2d`,
		},
		En: SDKLangInfo{
			Label:   `cocos2d-x`,
			DevTeam: `cocos2d`,
		},
	},
	{
		Soname: `librenpython.so`,
		Zh: SDKLangInfo{
			Label:   `Ren'Py`,
			DevTeam: `renpy.org`,
		},
		En: SDKLangInfo{
			Label:   `Ren'Py`,
			DevTeam: `renpy.org`,
		},
	},
	{
		Soname: `libfacedevice.so`,
		Zh: SDKLangInfo{
			Label:   `金融级实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Financial-grade Real-person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libx3g.so`,
		Zh: SDKLangInfo{
			Label:   `TopSec Reinforcement`,
			DevTeam: `顶象科技`,
		},
		En: SDKLangInfo{
			Label:   `顶象加固`,
			DevTeam: `顶象科技`,
		},
	},
	{
		Soname: `libreact_codegen_rngesturehandler_codegen.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libIMETable.so`,
		Zh: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `libIME`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libEasyAR.so`,
		Zh: SDKLangInfo{
			Label:   `EasyAR`,
			DevTeam: `视辰信息科技（上海）有限公司`,
		},
		En: SDKLangInfo{
			Label:   `EasyAR`,
			DevTeam: `视辰信息科技（上海）有限公司`,
		},
	},
	{
		Soname: `libreact_codegen_PrnkitSpec.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `liblbs.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯浏览服务（TBS）`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Browser Service (TBS)`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbarhopper.so`,
		Zh: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `BarHopper`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libuc_bhook_jni.so`,
		Zh: SDKLangInfo{
			Label:   `BHook`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BHook`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libuptsmservice.so`,
		Zh: SDKLangInfo{
			Label:   `银联 SDK`,
			DevTeam: `银联`,
		},
		En: SDKLangInfo{
			Label:   `UnionPay SDK`,
			DevTeam: `银联`,
		},
	},
	{
		Soname: `libaudio360.so`,
		Zh: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads Ltd`,
		},
		En: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads Ltd`,
		},
	},
	{
		Soname: `libmapbufferjni.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libimebra_lib.so`,
		Zh: SDKLangInfo{
			Label:   `imebra`,
			DevTeam: `imebra.com`,
		},
		En: SDKLangInfo{
			Label:   `imebra`,
			DevTeam: `imebra.com`,
		},
	},
	{
		Soname: `libEGL_angle.so`,
		Zh: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `angle`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libyuv_android.so`,
		Zh: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `libYUV`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libjsc.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libAVAPIsT.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libBugly-webank.so`,
		Zh: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `TDS 腾讯端服务`,
		},
		En: SDKLangInfo{
			Label:   `Bugly`,
			DevTeam: `Tencent Device-oriented Service`,
		},
	},
	{
		Soname: `libwcdb_legacy.so`,
		Zh: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `WCDB`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libmediacore.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `libchaquopy_java.so`,
		Zh: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
		En: SDKLangInfo{
			Label:   `Chaquopy`,
			DevTeam: `Chaquo Ltd`,
		},
	},
	{
		Soname: `libpng.so`,
		Zh: SDKLangInfo{
			Label:   `libpng`,
			DevTeam: `Greg Roelofs`,
		},
		En: SDKLangInfo{
			Label:   `libpng`,
			DevTeam: `Greg Roelofs`,
		},
	},
	{
		Soname: `libmnnqjs.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libmono-btls-shared.so`,
		Zh: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Xamarin`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libMuPDF.so`,
		Zh: SDKLangInfo{
			Label:   `MuPDF`,
			DevTeam: `Artifex Software Inc.`,
		},
		En: SDKLangInfo{
			Label:   `MuPDF`,
			DevTeam: `Artifex Software Inc.`,
		},
	},
	{
		Soname: `libAGFX.so`,
		Zh: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BytePlus RTC`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libhiai_model_compatible.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libzipalign.so`,
		Zh: SDKLangInfo{
			Label:   `zipalign`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `zipalign`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnativeaudioe7.so`,
		Zh: SDKLangInfo{
			Label:   `NativeAudio`,
			DevTeam: `Unity Technologies`,
		},
		En: SDKLangInfo{
			Label:   `NativeAudio`,
			DevTeam: `Unity Technologies`,
		},
	},
	{
		Soname: `libjsi.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libmediautil_v7.so`,
		Zh: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `AnyChat`,
			DevTeam: `广州佰锐网络科技有限公司`,
		},
	},
	{
		Soname: `liblposix.so`,
		Zh: SDKLangInfo{
			Label:   `ZAndrolua`,
			DevTeam: `MGLSIDE`,
		},
		En: SDKLangInfo{
			Label:   `ZAndrolua`,
			DevTeam: `MGLSIDE`,
		},
	},
	{
		Soname: `libscel2org5.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libquickjs-android.so`,
		Zh: SDKLangInfo{
			Label:   `QuickJS Binding`,
			DevTeam: `Harlon Wang`,
		},
		En: SDKLangInfo{
			Label:   `QuickJS Binding`,
			DevTeam: `Harlon Wang`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.lu.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libtquic_jni.so`,
		Zh: SDKLangInfo{
			Label:   `TQUIC`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `TQUIC`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libsymphony-cpu.so`,
		Zh: SDKLangInfo{
			Label:   `Symphony CPU 运行时库`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `Symphony CPU runtime library`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libmodelcrypto.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libUE4Bundle-Splittables.so`,
		Zh: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
		En: SDKLangInfo{
			Label:   `Unreal Engine`,
			DevTeam: `Epic Games`,
		},
	},
	{
		Soname: `libbasec_x86.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA Reinforcement`,
			DevTeam: `CFCA`,
		},
		En: SDKLangInfo{
			Label:   `CFCA 加固`,
			DevTeam: `CFCA`,
		},
	},
	{
		Soname: `libgmesdk.so`,
		Zh: SDKLangInfo{
			Label:   `GME`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `GME`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libbrotli_native.so`,
		Zh: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libogg_opus_encoder.so`,
		Zh: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
		En: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
	},
	{
		Soname: `libarcore_sdk_c.so`,
		Zh: SDKLangInfo{
			Label:   `Virocore`,
			DevTeam: `Viro Media`,
		},
		En: SDKLangInfo{
			Label:   `Virocore`,
			DevTeam: `Viro Media`,
		},
	},
	{
		Soname: `libusb.so`,
		Zh: SDKLangInfo{
			Label:   `libusb`,
			DevTeam: `libusb`,
		},
		En: SDKLangInfo{
			Label:   `libusb`,
			DevTeam: `libusb`,
		},
	},
	{
		Soname: `libAPMUOCPLIB.so`,
		Zh: SDKLangInfo{
			Label:   `APMInsight / 应用性能监控全链路版`,
			DevTeam: `Volcengine (火山引擎)`,
		},
		En: SDKLangInfo{
			Label:   `APMInsight / Full-link Application Performance Monitoring`,
			DevTeam: `Volcengine (火山引擎)`,
		},
	},
	{
		Soname: `libandfix.so`,
		Zh: SDKLangInfo{
			Label:   `AndFix`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `AndFix`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libwechatcrash.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libPixVideoCore.so`,
		Zh: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `PixUI`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtersafe.so`,
		Zh: SDKLangInfo{
			Label:   `手游安全 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Mobile Game Security SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libpanorenderer.so`,
		Zh: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Cardboard SDK`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libarcsoft_pet_detection.so`,
		Zh: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
		En: SDKLangInfo{
			Label:   `Arcsoft`,
			DevTeam: `Arcsoft`,
		},
	},
	{
		Soname: `libyoume_voice_engine.so`,
		Zh: SDKLangInfo{
			Label:   `游密`,
			DevTeam: `Youme`,
		},
		En: SDKLangInfo{
			Label:   `Youmi`,
			DevTeam: `Youme`,
		},
	},
	{
		Soname: `libxcrash.so`,
		Zh: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
		En: SDKLangInfo{
			Label:   `xCrash`,
			DevTeam: `iQIYI`,
		},
	},
	{
		Soname: `libxnn_zip.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libbytehook.so`,
		Zh: SDKLangInfo{
			Label:   `BHook`,
			DevTeam: `ByteDance`,
		},
		En: SDKLangInfo{
			Label:   `BHook`,
			DevTeam: `ByteDance`,
		},
	},
	{
		Soname: `libfree-reflection.so`,
		Zh: SDKLangInfo{
			Label:   `FreeReflection`,
			DevTeam: `tiann`,
		},
		En: SDKLangInfo{
			Label:   `FreeReflection`,
			DevTeam: `tiann`,
		},
	},
	{
		Soname: `libMNN_Vulkan.so`,
		Zh: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `MNN`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libnesdk.so`,
		Zh: SDKLangInfo{
			Label:   `网易云信`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `NetEase Cloud Letter`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `liblite_pybind.so`,
		Zh: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Paddle Lite`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libVisionCamera.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Vision Camera`,
			DevTeam: `mrousavy`,
		},
		En: SDKLangInfo{
			Label:   `React Native Vision Camera`,
			DevTeam: `mrousavy`,
		},
	},
	{
		Soname: `libdexkit.so`,
		Zh: SDKLangInfo{
			Label:   `DexKit`,
			DevTeam: `LuckyPray`,
		},
		En: SDKLangInfo{
			Label:   `DexKit`,
			DevTeam: `LuckyPray`,
		},
	},
	{
		Soname: `liblocationlocalenc.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯地图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Map SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libroute-tracepath.so`,
		Zh: SDKLangInfo{
			Label:   `tracepath`,
			DevTeam: `Netease`,
		},
		En: SDKLangInfo{
			Label:   `tracepath`,
			DevTeam: `Netease`,
		},
	},
	{
		Soname: `libssl.so`,
		Zh: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
		En: SDKLangInfo{
			Label:   `OpenSSL`,
			DevTeam: `OpenSSL`,
		},
	},
	{
		Soname: `libffmpeg.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libqmp.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯灯塔终端 ID 精准识别体系`,
			DevTeam: `腾讯灯塔`,
		},
		En: SDKLangInfo{
			Label:   `QIMEI SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libCNamaSDK.so`,
		Zh: SDKLangInfo{
			Label:   `Faceunity Nama SDK`,
			DevTeam: `杭州相芯科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Faceunity Nama SDK`,
			DevTeam: `杭州相芯科技有限公司`,
		},
	},
	{
		Soname: `libsnpe_loader.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libALBiometricsJni.so`,
		Zh: SDKLangInfo{
			Label:   `实人认证 SDK`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Real Person Authentication SDK`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsharewind.so`,
		Zh: SDKLangInfo{
			Label:   `微博 SDK`,
			DevTeam: `Weibo`,
		},
		En: SDKLangInfo{
			Label:   `Weibo SDK`,
			DevTeam: `Weibo`,
		},
	},
	{
		Soname: `libgifimage.so`,
		Zh: SDKLangInfo{
			Label:   `GIFLIB`,
			DevTeam: `GIFLIB`,
		},
		En: SDKLangInfo{
			Label:   `GIFLIB`,
			DevTeam: `GIFLIB`,
		},
	},
	{
		Soname: `libtool-checker.so`,
		Zh: SDKLangInfo{
			Label:   `RootBeer`,
			DevTeam: `Scott Alexander-Bown`,
		},
		En: SDKLangInfo{
			Label:   `RootBeer`,
			DevTeam: `Scott Alexander-Bown`,
		},
	},
	{
		Soname: `libemulator_check.so`,
		Zh: SDKLangInfo{
			Label:   `AntiFakerAndroidChecker`,
			DevTeam: `happylishang`,
		},
		En: SDKLangInfo{
			Label:   `AntiFakerAndroidChecker`,
			DevTeam: `happylishang`,
		},
	},
	{
		Soname: `libtoolChecker.so`,
		Zh: SDKLangInfo{
			Label:   `RootBeer`,
			DevTeam: `Scott Alexander-Bown`,
		},
		En: SDKLangInfo{
			Label:   `RootBeer`,
			DevTeam: `Scott Alexander-Bown`,
		},
	},
	{
		Soname: `libsqlcipher.so`,
		Zh: SDKLangInfo{
			Label:   `SQLCipher`,
			DevTeam: `Zetetic`,
		},
		En: SDKLangInfo{
			Label:   `SQLCipher`,
			DevTeam: `Zetetic`,
		},
	},
	{
		Soname: `libavdevice.so`,
		Zh: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
		En: SDKLangInfo{
			Label:   `FFmpeg`,
			DevTeam: `FFmpeg`,
		},
	},
	{
		Soname: `libMicrosoft.CognitiveServices.Speech.extension.codec.so`,
		Zh: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
		En: SDKLangInfo{
			Label:   `Cognitive Services Speech SDK`,
			DevTeam: `Microsoft`,
		},
	},
	{
		Soname: `libjpgt.so`,
		Zh: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
		En: SDKLangInfo{
			Label:   `Leptonica`,
			DevTeam: `Dan Bloomberg`,
		},
	},
	{
		Soname: `libandroidx.graphics.path.so`,
		Zh: SDKLangInfo{
			Label:   `Jetpack Graphics`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Jetpack Graphics`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libpingpp.so`,
		Zh: SDKLangInfo{
			Label:   `Ping++ SDK`,
			DevTeam: `Ping++`,
		},
		En: SDKLangInfo{
			Label:   `Ping++ SDK`,
			DevTeam: `Ping++`,
		},
	},
	{
		Soname: `libsense.so`,
		Zh: SDKLangInfo{
			Label:   `libsense`,
			DevTeam: `moshigottlieb`,
		},
		En: SDKLangInfo{
			Label:   `libsense`,
			DevTeam: `moshigottlieb`,
		},
	},
	{
		Soname: `libsnpe_dsp_domains_v2_system.so`,
		Zh: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
		En: SDKLangInfo{
			Label:   `SNPE SDK`,
			DevTeam: `Qualcomm Snapdragon`,
		},
	},
	{
		Soname: `libbasec.so`,
		Zh: SDKLangInfo{
			Label:   `CFCA Reinforcement`,
			DevTeam: `CFCA`,
		},
		En: SDKLangInfo{
			Label:   `CFCA 加固`,
			DevTeam: `CFCA`,
		},
	},
	{
		Soname: `libRSSupport.so`,
		Zh: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
		En: SDKLangInfo{
			Label:   `RenderScript`,
			DevTeam: `Android`,
		},
	},
	{
		Soname: `libadmob.so`,
		Zh: SDKLangInfo{
			Label:   `Google AdMob`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Google AdMob`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `liblimbo.so`,
		Zh: SDKLangInfo{
			Label:   `Limbo`,
			DevTeam: `limboemu`,
		},
		En: SDKLangInfo{
			Label:   `Limbo`,
			DevTeam: `limboemu`,
		},
	},
	{
		Soname: `libjiagu_vip.so`,
		Zh: SDKLangInfo{
			Label:   `360 加固`,
			DevTeam: `360`,
		},
		En: SDKLangInfo{
			Label:   `360 Reinforcement`,
			DevTeam: `360`,
		},
	},
	{
		Soname: `libsdl.so`,
		Zh: SDKLangInfo{
			Label:   `libSDL`,
			DevTeam: `libsdl.org`,
		},
		En: SDKLangInfo{
			Label:   `libSDL`,
			DevTeam: `libsdl.org`,
		},
	},
	{
		Soname: `libYTCommonXMagic.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯优图 SDK`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Youtu SDK`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libexecoat_x86.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libfreetype-jni.so`,
		Zh: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
		En: SDKLangInfo{
			Label:   `FreeType`,
			DevTeam: `FreeType`,
		},
	},
	{
		Soname: `libbaiduprotect_sec_jni_facect.so`,
		Zh: SDKLangInfo{
			Label:   `百度应用加固`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu App Reinforcement`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libmatrix-jectl.so`,
		Zh: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Matrix`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libxquic.so`,
		Zh: SDKLangInfo{
			Label:   `XQUIC`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `XQUIC`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `librnscreens.so`,
		Zh: SDKLangInfo{
			Label:   `React Native Screens`,
			DevTeam: `Software Mansion`,
		},
		En: SDKLangInfo{
			Label:   `React Native Screens`,
			DevTeam: `Software Mansion`,
		},
	},
	{
		Soname: `libkraken_jsc.so`,
		Zh: SDKLangInfo{
			Label:   `Kraken`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `Kraken`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libnmmp.so`,
		Zh: SDKLangInfo{
			Label:   `nmmp`,
			DevTeam: `maoabc`,
		},
		En: SDKLangInfo{
			Label:   `nmmp`,
			DevTeam: `maoabc`,
		},
	},
	{
		Soname: `libsentry-fd-monitor.so.so`,
		Zh: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
		En: SDKLangInfo{
			Label:   `Sentry`,
			DevTeam: `Sentry`,
		},
	},
	{
		Soname: `libchttrans.so`,
		Zh: SDKLangInfo{
			Label:   `Fcitx 5 中文插件`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
		En: SDKLangInfo{
			Label:   `Fcitx 5 Chinese Addons`,
			DevTeam: `Fcitx / Fcitx 5 for Android`,
		},
	},
	{
		Soname: `libogg_opus_encoder_translate.so`,
		Zh: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
		En: SDKLangInfo{
			Label:   `Opus`,
			DevTeam: `Xiph.Org`,
		},
	},
	{
		Soname: `libfbjni.so`,
		Zh: SDKLangInfo{
			Label:   `Facebook SDK`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Facebook SDK`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libAVAPIs.so`,
		Zh: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
		En: SDKLangInfo{
			Label:   `TUTK SDK`,
			DevTeam: `TUTK`,
		},
	},
	{
		Soname: `libgetuiext2.so`,
		Zh: SDKLangInfo{
			Label:   `个推`,
			DevTeam: `个推`,
		},
		En: SDKLangInfo{
			Label:   `Getui`,
			DevTeam: `个推`,
		},
	},
	{
		Soname: `libnative-filters.so`,
		Zh: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Fresco`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `liblbsdl.so`,
		Zh: SDKLangInfo{
			Label:   `乐播 SDK`,
			DevTeam: `Hpplay`,
		},
		En: SDKLangInfo{
			Label:   `Lebo SDK`,
			DevTeam: `Hpplay`,
		},
	},
	{
		Soname: `libucrop.so`,
		Zh: SDKLangInfo{
			Label:   `uCrop`,
			DevTeam: `Yalantis`,
		},
		En: SDKLangInfo{
			Label:   `uCrop`,
			DevTeam: `Yalantis`,
		},
	},
	{
		Soname: `libliantian.so`,
		Zh: SDKLangInfo{
			Label:   `百度人脸识别`,
			DevTeam: `Baidu`,
		},
		En: SDKLangInfo{
			Label:   `Baidu Face Recognition`,
			DevTeam: `Baidu`,
		},
	},
	{
		Soname: `libandroidndkgif.so`,
		Zh: SDKLangInfo{
			Label:   `Android NDK GIF`,
			DevTeam: `waynejo`,
		},
		En: SDKLangInfo{
			Label:   `Android NDK GIF`,
			DevTeam: `waynejo`,
		},
	},
	{
		Soname: `libreact_newarchdefaults.so`,
		Zh: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `React Native`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `liblibusb.so`,
		Zh: SDKLangInfo{
			Label:   `libusb`,
			DevTeam: `libusb`,
		},
		En: SDKLangInfo{
			Label:   `libusb`,
			DevTeam: `libusb`,
		},
	},
	{
		Soname: `libbrotlicommon.so`,
		Zh: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Brotli`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libnavermap.so`,
		Zh: SDKLangInfo{
			Label:   `Naver Map`,
			DevTeam: `LBSTECH`,
		},
		En: SDKLangInfo{
			Label:   `Naver Map`,
			DevTeam: `LBSTECH`,
		},
	},
	{
		Soname: `libAVProLocal.so`,
		Zh: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads`,
		},
		En: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads`,
		},
	},
	{
		Soname: `libiconv16.so`,
		Zh: SDKLangInfo{
			Label:   `iconv`,
			DevTeam: `GNU`,
		},
		En: SDKLangInfo{
			Label:   `iconv`,
			DevTeam: `GNU`,
		},
	},
	{
		Soname: `libhermes-executor-common-release.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libmaplibre.so`,
		Zh: SDKLangInfo{
			Label:   `MapLibre Native`,
			DevTeam: `The MapLibre Organization`,
		},
		En: SDKLangInfo{
			Label:   `MapLibre Native`,
			DevTeam: `The MapLibre Organization`,
		},
	},
	{
		Soname: `libvosk.so`,
		Zh: SDKLangInfo{
			Label:   `Vosk`,
			DevTeam: `alphacephei`,
		},
		En: SDKLangInfo{
			Label:   `Vosk`,
			DevTeam: `alphacephei`,
		},
	},
	{
		Soname: `libhermestooling.so`,
		Zh: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
		En: SDKLangInfo{
			Label:   `Hermes JS Engine`,
			DevTeam: `Facebook`,
		},
	},
	{
		Soname: `libaudio360-JNI.so`,
		Zh: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads Ltd`,
		},
		En: SDKLangInfo{
			Label:   `AVPro Video`,
			DevTeam: `RenderHeads Ltd`,
		},
	},
	{
		Soname: `libSparkChain.so`,
		Zh: SDKLangInfo{
			Label:   `星火认知大模型`,
			DevTeam: `科大讯飞`,
		},
		En: SDKLangInfo{
			Label:   `Sparks Cognitive Large Model`,
			DevTeam: `科大讯飞`,
		},
	},
	{
		Soname: `libmtc.so`,
		Zh: SDKLangInfo{
			Label:   `JusTalk Cloud`,
			DevTeam: `宁波菊风系统软件有限公司`,
		},
		En: SDKLangInfo{
			Label:   `JusTalk Cloud`,
			DevTeam: `宁波菊风系统软件有限公司`,
		},
	},
	{
		Soname: `libhiai_v1cl.so`,
		Zh: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HiAI Foundation`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libjniPdfium.so`,
		Zh: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
		En: SDKLangInfo{
			Label:   `Pdfium`,
			DevTeam: `Google`,
		},
	},
	{
		Soname: `libtanx.so`,
		Zh: SDKLangInfo{
			Label:   `TanX`,
			DevTeam: `Alibaba`,
		},
		En: SDKLangInfo{
			Label:   `TanX`,
			DevTeam: `Alibaba`,
		},
	},
	{
		Soname: `libsandhook-art.so`,
		Zh: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
		En: SDKLangInfo{
			Label:   `Sandhook`,
			DevTeam: `asLody`,
		},
	},
	{
		Soname: `libreact_codegen_rnsvg.so`,
		Zh: SDKLangInfo{
			Label:   `react-native-svg`,
			DevTeam: `software-mansion`,
		},
		En: SDKLangInfo{
			Label:   `react-native-svg`,
			DevTeam: `software-mansion`,
		},
	},
	{
		Soname: `libdcblur.so`,
		Zh: SDKLangInfo{
			Label:   `DCloud`,
			DevTeam: `数字天堂`,
		},
		En: SDKLangInfo{
			Label:   `DCloud`,
			DevTeam: `数字天堂`,
		},
	},
	{
		Soname: `libagora-fdkaac.so`,
		Zh: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
		En: SDKLangInfo{
			Label:   `Agora RTC SDK`,
			DevTeam: `Agora`,
		},
	},
	{
		Soname: `libhdr_pro_engine.so`,
		Zh: SDKLangInfo{
			Label:   `HMS Video Kit HDR Vivid`,
			DevTeam: `HUAWEI`,
		},
		En: SDKLangInfo{
			Label:   `HMS Video Kit HDR Vivid`,
			DevTeam: `HUAWEI`,
		},
	},
	{
		Soname: `libxul.so`,
		Zh: SDKLangInfo{
			Label:   `Lib_XUL`,
			DevTeam: `Mozilla`,
		},
		En: SDKLangInfo{
			Label:   `Lib_XUL`,
			DevTeam: `Mozilla`,
		},
	},
	{
		Soname: `liblexbor.so`,
		Zh: SDKLangInfo{
			Label:   `Lexbor`,
			DevTeam: `Alexander Borisov`,
		},
		En: SDKLangInfo{
			Label:   `Lexbor`,
			DevTeam: `Alexander Borisov`,
		},
	},
	{
		Soname: `libgdx-box2d.so`,
		Zh: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
		En: SDKLangInfo{
			Label:   `libGDX`,
			DevTeam: `libgdx`,
		},
	},
	{
		Soname: `libconnectionbase.so`,
		Zh: SDKLangInfo{
			Label:   `小米游戏 SDK`,
			DevTeam: `Xiaomi`,
		},
		En: SDKLangInfo{
			Label:   `Xiaomi Game SDK`,
			DevTeam: `Xiaomi`,
		},
	},
	{
		Soname: `libexec.so`,
		Zh: SDKLangInfo{
			Label:   `爱加密`,
			DevTeam: `北京智游网安科技有限公司`,
		},
		En: SDKLangInfo{
			Label:   `Aijia Encryption`,
			DevTeam: `北京智游网安科技有限公司`,
		},
	},
	{
		Soname: `libzlib.so`,
		Zh: SDKLangInfo{
			Label:   `zlib`,
			DevTeam: `zlib`,
		},
		En: SDKLangInfo{
			Label:   `zlib`,
			DevTeam: `zlib`,
		},
	},
	{
		Soname: `libunrar-jni.so`,
		Zh: SDKLangInfo{
			Label:   `unrar-android`,
			DevTeam: `maoabc`,
		},
		En: SDKLangInfo{
			Label:   `unrar-android`,
			DevTeam: `maoabc`,
		},
	},
	{
		Soname: `libopenh264.so`,
		Zh: SDKLangInfo{
			Label:   `OpenH264`,
			DevTeam: `cisco`,
		},
		En: SDKLangInfo{
			Label:   `OpenH264`,
			DevTeam: `cisco`,
		},
	},
	{
		Soname: `libxnn_media.so`,
		Zh: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
		En: SDKLangInfo{
			Label:   `xnn`,
			DevTeam: `aaalgo`,
		},
	},
	{
		Soname: `libmtanativecrash_v2.so`,
		Zh: SDKLangInfo{
			Label:   `腾讯移动分析`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Tencent Mobile Analytics`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libwxaopenruntimejni.so`,
		Zh: SDKLangInfo{
			Label:   `Donut`,
			DevTeam: `Tencent`,
		},
		En: SDKLangInfo{
			Label:   `Donut`,
			DevTeam: `Tencent`,
		},
	},
	{
		Soname: `libtun2proxy.so`,
		Zh: SDKLangInfo{
			Label:   `tun2proxy`,
			DevTeam: `tun2proxy`,
		},
		En: SDKLangInfo{
			Label:   `tun2proxy`,
			DevTeam: `tun2proxy`,
		},
	},
}
