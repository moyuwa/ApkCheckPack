#!/usr/bin/env python3
# coding=utf-8
# python version 3.7 by 6time
#

import zipfile, hashlib
from tkinter import *
from tkinter.filedialog import askdirectory
import windnd
from tkinter.messagebox import showinfo

markNameMap = {}
markNameMap["libchaosvmp.so"] = "娜迦";
markNameMap["libddog.so"] = "娜迦"
markNameMap["libfdog.so"] = "娜迦"
markNameMap["libedog.so"] = "娜迦企业版"
markNameMap["libexec.so"] = "爱加密"
markNameMap["libexecmain.so"] = "爱加密"
markNameMap["ijiami.dat"] = "爱加密"
markNameMap["ijiami.ajm"] = "爱加密企业版"
markNameMap["libsecexe.so"] = "梆梆免费版"
markNameMap["libsecmain.so"] = "梆梆免费版"
markNameMap["libSecShell.so"] = "梆梆免费版"
markNameMap["libDexHelper.so"] = "梆梆企业版"
markNameMap["libDexHelper-x86.so"] = "梆梆企业版"
markNameMap["libprotectClass.so"] = "360"
markNameMap["libjiagu.so"] = "360"
markNameMap["libjiagu_art.so"] = "360"
markNameMap["libjiagu_x86.so"] = "360"
markNameMap["libegis.so"] = "通付盾"
markNameMap["libNSaferOnly.so"] = "通付盾"
markNameMap["libnqshield.so"] = "网秦"
markNameMap["libbaiduprotect.so"] = "百度"
markNameMap["aliprotect.dat"] = "阿里聚安全"
markNameMap["libsgmain.so"] = "阿里聚安全"
markNameMap["libsgsecuritybody.so"] = "阿里聚安全"
markNameMap["libmobisec.so"] = "阿里聚安全"
markNameMap["libtup.so"] = "腾讯"
markNameMap["libexec.so"] = "腾讯"
markNameMap["libshell.so"] = "腾讯"
markNameMap["mix.dex"] = "腾讯"
markNameMap["lib/armeabi/mix.dex"] = "腾讯"
markNameMap["lib/armeabi/mixz.dex"] = "腾讯"
markNameMap["tencent_sub"] = "腾讯云"
markNameMap["lib/armeabi/libshell-super.2019.so"] = "腾讯云"
markNameMap["lib/armeabi/libshell-super.2020.so"] = "腾讯云"
markNameMap["lib/armeabi/libshell-super.2021.so"] = "腾讯云"
markNameMap["assets/libshellx-super.2021.so"] = "腾讯云"
markNameMap["lib/armeabi-v7a/libxgVipSecurity.so"] = "腾讯乐固（VMP）"
markNameMap["lib/arm64-v8a/libxgVipSecurity.so"] = "腾讯乐固（VMP）"
markNameMap["libtosprotection.armeabi.so"] = "腾讯御安全"
markNameMap["libtosprotection.armeabi-v7a.so"] = "腾讯御安全"
markNameMap["libtosprotection.x86.so"] = "腾讯御安全"
markNameMap["libnesec.so"] = "网易易盾"
markNameMap["libAPKProtect.so"] = "APKProtect"
markNameMap["libkwscmm.so"] = "几维安全"
markNameMap["libkwscr.so"] = "几维安全"
markNameMap["libkwslinker.so"] = "几维安全"
markNameMap["libx3g.so"] = "顶像科技"
markNameMap["libapssec.so"] = "盛大"
markNameMap["librsprotect.so"] = "瑞星"
markNameMap["assets/mxsafe.jar"] = "蛮犀"
markNameMap["assets/mxsafe.config"] = "蛮犀"
markNameMap["assets/mxsafe.data"] = "蛮犀"
markNameMap["assets/mxsafe/arm64-v8a/libdSafeShell.so"] = "蛮犀"
markNameMap["assets/mxsafe/x86_64/libdSafeShell.so"] = "蛮犀"


def check_jiagu(filename):
    azip = zipfile.ZipFile(filename)  # 默认模式r,读
    # print(azip.namelist()) # 返回所有文件夹和文件
    # print(azip.filename)# 返回该zip的文件名
    for zippath in azip.namelist():
        if 'lib' in zippath:
            # print(zippath)
            for mark in markNameMap.keys():
                if mark in zippath:
                    print(u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(markNameMap[mark], zippath, mark))
                    return (u"检测到 【{}】 加固\n匹配特征:{}\n->{}\n".format(markNameMap[mark], zippath, mark))
    # so库文件模式找不到就全量匹配
    for zippath in azip.namelist():
        for mark in markNameMap.keys():
            if mark in zippath:
                print(u"检测到 【{}】 加固\n匹配特征:{}->{}\n".format(markNameMap[mark], zippath, mark))
                return (u"【全量匹配】检测到 【{}】 加固\n匹配特征:{}\n->{}\n".format(markNameMap[mark], zippath, mark))

    return (u"未检测到加固")


def md5sum(file_name):
    with open(file_name, 'rb') as fp:
        data = fp.read()
    file_md5 = hashlib.md5(data).hexdigest()
    print(file_md5)
    pathmd5.set(file_md5)


def selectPath():
    path_ = askdirectory()
    path.set(path_)


def dragged_files(files):
    msg = '\n'.join((item.decode('gbk') for item in files))
    path.set(msg)
    print(msg)
    if msg[-4:] != '.apk':
        pathjiagu.set("非APK文件！")
    else:
        _jiagu = check_jiagu(msg)
        pathjiagu.set(_jiagu)
        md5sum(msg)


if __name__ == '__main__':
    print(u'基于 "ApkScan-PKID查壳工具" 的规则 添加\修改 by 6time\n\n\n')
    help = ("""目前支持的加固厂商：
        娜迦、梆梆、爱加密、通付盾、360加固、
        百度加固、阿里加固、腾讯加固、 网易易盾
        盛大加固、瑞星加固、网秦加固、
        国信灵通加固、几维安全加固、顶像科技加固、
        apkprotect加固、等常规厂商
用法:        
    python3 apkscan-pkid.py D:\\android\\app.apk
        """)
    # if len(sys.argv) != 2:
    #     print(help)
    # else:
    #     check_jiagu(sys.argv[1])  # sys.argv[1]

    root = Tk()
    root.title("ApkScan GUI v2.1 by 6time")
    # 获取屏幕 宽、高
    ws = root.winfo_screenwidth()
    hs = root.winfo_screenheight()
    root.geometry('%dx%d+%d+%d' % (500, 300, ws / 2 - 100, hs / 2 - 100))
    root.resizable(False, False)  # 窗口大小不可变
    path = StringVar()
    pathjiagu = StringVar()
    pathmd5 = StringVar()
    # Label(root, text="目标路径:").grid(row=0, column=0)
    # Entry(root, textvariable=path).grid(row=0, column=1)
    # Button(root, text="路径选择", command=selectPath).grid(row=0, column=2)
    # Label(root, text="目标路径:").grid(row=1, column=0)
    Label(root, text="APK文件路径").pack(fill=BOTH)
    Entry(root, textvariable=path).pack(fill=BOTH)
    windnd.hook_dropfiles(root, func=dragged_files)
    # Button(root, text="路径选择", command=selectPath).pack(fill=BOTH)
    Label(root, text="检测结果").pack(fill=BOTH)
    Label(root, textvariable=pathjiagu).pack(fill=BOTH)
    Label(root, text="文件MD5值").pack(fill=BOTH)
    Label(root, textvariable=pathmd5).pack(fill=BOTH)
    root.mainloop()
