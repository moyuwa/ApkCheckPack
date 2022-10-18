#!/usr/bin/env python3
# coding=utf-8
# python version 3.7 by 6time
#

import zipfile, hashlib, json
from tkinter import *
from tkinter.filedialog import askdirectory, askopenfilename
from tkinter.messagebox import showinfo
import windnd

packs = """{
  "360加固": [
    "assets/.appkey",
    "assets/libjiagu.so",
    "libjiagu.so",
    "libjiagu_art.so",
    "libjiagu_x86.so",
    "libprotectClass.so",
    ".appkey",
    "1ibjgdtc.so",
    "libjgdtc.so",
    "libjgdtc_a64.so",
    "libjgdtc_art.so",
    "libjgdtc_x64.so",
    "libjgdtc_x86.so",
    "libjiagu_a64.so",
    "libjiagu_ls.so",
    "libjiagu_x64.so"
  ],
  "APKProtect": [
    "libAPKProtect.so"
  ],
  "UU安全": [
    "libuusafe.jar.so",
    "libuusafe.so",
    "libuusafeempty.so",
    "assets/libuusafe.jar.so",
    "assets/libuusafe.so",
    "lib/armeabi/libuusafeempty.so"
  ],
  "apktoolplus": [
    "assets/jiagu_data.bin",
    "assets/sign.bin",
    "jiagu_data.bin",
    "lib/armeabi/libapktoolplus_jiagu.so",
    "libapktoolplus_jiagu.so",
    "sign.bin"
  ],
  "中国移动加固": [
    "assets/mogosec_classes",
    "assets/mogosec_data",
    "assets/mogosec_dexinfo",
    "assets/mogosec_march",
    "ibmogosecurity.so",
    "lib/armeabi/libcmvmp.so",
    "lib/armeabi/libmogosec_dex.so",
    "lib/armeabi/libmogosec_sodecrypt.so",
    "lib/armeabi/libmogosecurity.so",
    "libcmvmp.so",
    "libmogosec_dex.so",
    "libmogosec_sodecrypt.so",
    "mogosec_classes",
    "mogosec_data",
    "mogosec_dexinfo",
    "mogosec_march"
  ],
  "几维安全": [
    "assets/dex.dat",
    "lib/armeabi/kdpdata.so",
    "lib/armeabi/libkdp.so",
    "lib/armeabi/libkwscmm.so",
    "libkwscmm.so",
    "libkwscr.so",
    "libkwslinker.so"
  ],
  "启明星辰": [
    "libvenSec.so",
    "libvenustech.so"
  ],
  "网秦加固": [
    "libnqshield.so"
  ],
  "娜迦加固": [
    "libchaosvmp.so",
    "libddog.so",
    "libfdog.so"
  ],
  "娜迦加固（新版2022）": [
    "assets/maindata/fake_classes.dex",
    "lib/armeabi/libxloader.so",
    "lib/armeabi-v7a/libxloader.so",
    "lib/arm64-v8a/libxloader.so",
    "libxloader.so"
  ],
  "娜迦加固（企业版）": [
    "libedog.so"
  ],
  "梆梆安全（企业版）": [
    "libDexHelper-x86.so",
    "libDexHelper.so",
    "1ibDexHelper.so"
  ],
  "梆梆安全": [
    "libSecShell.so",
    "libsecexe.so",
    "libsecmain.so",
    "libSecShel1.so"
  ],
  "梆梆安全（定制版）": [
    "assets/classes.jar",
    "lib/armeabi/DexHelper.so"
  ],
  "梆梆安全（免费版）": [
    "assets/secData0.jar",
    "lib/armeabi/libSecShell-x86.so",
    "lib/armeabi/libSecShell.so"
  ],
  "海云安加固": [
    "assets/itse",
    "lib/armeabi/libitsec.so",
    "libitsec.so"
  ],
  "爱加密": [
    "assets/af.bin",
    "assets/ijiami.ajm",
    "assets/ijm_lib/X86/libexec.so",
    "assets/ijm_lib/armeabi/libexec.so",
    "assets/signed.bin",
    "ijiami.dat",
    "lib/armeabi/libexecmain.so",
    "libexecmain.so"
  ],
  "爱加密企业版": [
    "ijiami.ajm"
  ],
  "珊瑚灵御": [
    "assets/libreincp.so",
    "assets/libreincp_x86.so",
    "libreincp.so",
    "libreincp_x86.so"
  ],
  "瑞星加固": [
    "librsprotect.so"
  ],
  "百度加固": [
    "libbaiduprotect.so",
    "assets/baiduprotect.jar",
    "assets/baiduprotect1.jar",
    "baiduprotect1.jar",
    "lib/armeabi/libbaiduprotect.so",
    "libbaiduprotect_art.so",
    "libbaiduprotect_x86.so"
  ],
  "盛大加固": [
    "libapssec.so"
  ],
  "网易易盾": [
    "libnesec.so"
  ],
  "腾讯": [
    "libexec.so",
    "libshell.so"
  ],
  "腾讯加固": [
    "lib/armeabi/mix.dex",
    "lib/armeabi/mixz.dex",
    "lib/armeabi/libshella-xxxx.so",
    "lib/armeabi/libshellx-xxxx.so",
    "tencent_stub"
  ],
  "腾讯乐固（旧版）": [
    "libtup.so",
    "mix.dex",
    "liblegudb.so",
    "libshella",
    "mixz.dex",
    "libshel1x"
  ],
  "腾讯乐固": [
    "libshellx"
  ],
  "腾讯乐固（VMP）": [
    "lib/arm64-v8a/libxgVipSecurity.so",
    "lib/armeabi-v7a/libxgVipSecurity.so",
    "libxgVipSecurity.so"
  ],
  "腾讯云": [
    "assets/libshellx-super.2021.so",
    "lib/armeabi/libshell-super.2019.so",
    "lib/armeabi/libshell-super.2020.so",
    "lib/armeabi/libshell-super.2021.so",
    "lib/armeabi/libshell-super.2022.so",
    "lib/armeabi/libshell-super.2023.so",
    "tencent_sub"
  ],
  "腾讯云移动应用安全": [
    "0000000lllll.dex",
    "00000olllll.dex",
    "000O00ll111l.dex",
    "00O000ll111l.dex",
    "0OO00l111l1l",
    "o0oooOO0ooOo.dat"
  ],
  "腾讯云移动应用安全（腾讯御安全）": [
    "libBugly-yaq.so",
    "libshell-super.2019.so",
    "libshellx-super.2019.so",
    "libzBugly-yaq.so",
    "t86",
    "tosprotection",
    "tosversion",
    "000000011111.dex",
    "000000111111.dex",
    "000001111111",
    "00000o11111.dex",
    "o0ooo000oo0o.dat"
  ],
  "腾讯御安全": [
    "libtosprotection.armeabi-v7a.so",
    "libtosprotection.armeabi.so",
    "libtosprotection.x86.so",
    "assets/libtosprotection.armeabi-v7a.so",
    "assets/libtosprotection.armeabi.so",
    "assets/libtosprotection.x86.so",
    "assets/tosversion",
    "lib/armeabi/libTmsdk-xxx-mfr.so",
    "lib/armeabi/libtest.so"
  ],
  "腾讯Bugly": [
    "lib/arm64-v8a/libBugly.so",
    "libBugly.so"
  ],
  "蛮犀": [
    "assets/mxsafe.config",
    "assets/mxsafe.data",
    "assets/mxsafe.jar",
    "assets/mxsafe/arm64-v8a/libdSafeShell.so",
    "assets/mxsafe/x86_64/libdSafeShell.so",
    "libdSafeShell.so"
  ],
  "通付盾": [
    "libNSaferOnly.so",
    "libegis.so"
  ],
  "阿里加固": [
    "assets/armeabi/libfakejni.so",
    "assets/armeabi/libzuma.so",
    "assets/classes.dex.dat",
    "assets/dp.arm-v7.so.dat",
    "assets/dp.arm.so.dat",
    "assets/libpreverify1.so",
    "assets/libzuma.so",
    "assets/libzumadata.so",
    "dexprotect"
  ],
  "阿里聚安全": [
    "aliprotect.dat",
    "libdemolish.so",
    "libfakejni.so",
    "libmobisec.so",
    "libsgmain.so",
    "libzuma.so",
    "libzumadata.so",
    "libdemolishdata.so",
    "libpreverify1.so",
    "libsgsecuritybody.so"
  ],
  "顶像科技": [
    "libx3g.so",
    "lib/armeabi/libx3g.so"
  ]
}"""

# json文件方式加载特征
# with open('pack.json', 'r', encoding='utf-8') as f:
#     markNameMap = json.load(f)
#     markNameMap = dict(markNameMap)
# 字符串方式加载特征
markNameMap = json.loads(packs)
markNameMap = dict(markNameMap)


def check_jiagu(filename):
    azip = zipfile.ZipFile(filename)  # 默认模式r,读
    # print(azip.namelist()) # 返回所有文件夹和文件
    # print(azip.filename)# 返回该zip的文件名
    jigu = u''
    for zippath in azip.namelist():
        if 'lib' in zippath:
            # print(zippath)
            for key, value in markNameMap.items():
                for mark in value:
                    if mark in zippath:
                        print(u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(key, zippath, mark))
                        jigu += (u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(key, zippath, mark))
                        # return (u"检测到 【{}】 加固\n匹配特征:{}\n->{}\n".format(key, zippath, mark))
    if len(jigu) > 0:
        return jigu
    # so库文件模式找不到就全量匹配
    for zippath in azip.namelist():
        for key, value in markNameMap.items():
            for mark in value:
                if mark in zippath:
                    print(u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(key, zippath, mark))
                    # return (u"【全量匹配】检测到 【{}】 加固\n匹配特征:{}\n->{}\n".format(key, zippath, mark))
                    jigu += (u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(key, zippath, mark))
    if len(jigu) > 0:
        return jigu
    return (u"未检测到加固")


def md5sum(file_name):
    with open(file_name, 'rb') as fp:
        data = fp.read()
    file_md5 = hashlib.md5(data).hexdigest()
    print(file_md5)
    pathmd5.set(file_md5)


def selectfilePath():
    path_ = askopenfilename()
    path.set(path_)
    if path_[-4:] != '.apk':
        pathjiagu.set("非APK文件！")
    else:
        _jiagu = check_jiagu(path_)
        pathjiagu.set(_jiagu)
        md5sum(path_)

def dragged_files(files):
    msg = '\n'.join((item.decode('gbk') for item in files))
    print(msg)
    path.set(msg)
    if msg[-4:] != '.apk':
        pathjiagu.set("非APK文件！")
    else:
        _jiagu = check_jiagu(msg)
        pathjiagu.set(_jiagu)
        md5sum(msg)

if __name__ == '__main__':
    help = ("""ApkCheckPack v1.1 by 6time\n
    目前规则总数：177条
        """)
    root = Tk()
    root.title("ApkCheckPack v1.1 by 6time")
    # 获取屏幕 宽、高
    ws = root.winfo_screenwidth()
    hs = root.winfo_screenheight()
    root.geometry('%dx%d+%d+%d' % (600, 500, ws / 2, 0))
    root.resizable(True, True)  # 窗口大小可变
    path = StringVar()
    pathjiagu = StringVar()
    pathmd5 = StringVar()
    Label(root, text="APK文件路径").pack(fill=BOTH)
    Entry(root, textvariable=path).pack(fill=BOTH)

    windnd.hook_dropfiles(root, func=dragged_files)
    # Button(root, text="路径选择", command=selectfilePath).pack(fill=BOTH)

    Label(root, text="文件MD5值").pack(fill=BOTH)
    Label(root, textvariable=pathmd5).pack(fill=BOTH)
    Label(root, text="检测结果").pack(fill=BOTH)
    Label(root, textvariable=pathjiagu).pack(fill=BOTH)
    root.mainloop()
