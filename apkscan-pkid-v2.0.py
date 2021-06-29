#!/usr/bin/env python3
# coding=utf-8
# python version 3.7 by 6time
#
import sys
import zipfile

markNameMap = {}
markNameMap["libchaosvmp.so"]= "娜迦";
markNameMap["libddog.so"]= "娜迦"
markNameMap["libfdog.so"]= "娜迦"
markNameMap["libedog.so"]= "娜迦企业版"

markNameMap["libexec.so"]= "爱加密"
markNameMap["libexecmain.so"]= "爱加密"
markNameMap["ijiami.dat"]= "爱加密"
markNameMap["ijiami.ajm"]= "爱加密企业版"

markNameMap["libsecexe.so"]= "梆梆免费版"
markNameMap["libsecmain.so"]= "梆梆免费版"
markNameMap["libSecShell.so"]= "梆梆免费版"
markNameMap["libDexHelper.so"]= "梆梆企业版"
markNameMap["libDexHelper-x86.so"]= "梆梆企业版"

markNameMap["libprotectClass.so"]= "360"
markNameMap["libjiagu.so"]= "360"
markNameMap["libjiagu_art.so"]= "360"
markNameMap["libjiagu_x86.so"]= "360"

markNameMap["libegis.so"]= "通付盾"
markNameMap["libNSaferOnly.so"]= "通付盾"

markNameMap["libnqshield.so"]= "网秦"

markNameMap["libbaiduprotect.so"]= "百度"

markNameMap["aliprotect.dat"]= "阿里聚安全"
markNameMap["libsgmain.so"]= "阿里聚安全"
markNameMap["libsgsecuritybody.so"]= "阿里聚安全"
markNameMap["libmobisec.so"]= "阿里聚安全"

markNameMap["libtup.so"]= "腾讯"
markNameMap["libexec.so"]= "腾讯"
markNameMap["libshell.so"]= "腾讯"
markNameMap["mix.dex"]= "腾讯"
markNameMap["lib/armeabi/mix.dex"]= "腾讯"
markNameMap["lib/armeabi/mixz.dex"]= "腾讯"

markNameMap["tencent_sub"]="腾讯云"
markNameMap["lib/armeabi/libshell-super.2019.so"]="腾讯云"
markNameMap["lib/armeabi/libshell-super.2020.so"]= "腾讯云"
markNameMap["lib/armeabi/libshell-super.2021.so"]="腾讯云"
markNameMap["assets/libshellx-super.2021.so"]="腾讯云"

markNameMap["libtosprotection.armeabi.so"]= "腾讯御安全"
markNameMap["libtosprotection.armeabi-v7a.so"]= "腾讯御安全"
markNameMap["libtosprotection.x86.so"]= "腾讯御安全"

markNameMap["libnesec.so"]= "网易易盾"

markNameMap["libAPKProtect.so"]= "APKProtect"

markNameMap["libkwscmm.so"]= "几维安全"
markNameMap["libkwscr.so"]= "几维安全"
markNameMap["libkwslinker.so"]= "几维安全"

markNameMap["libx3g.so"]= "顶像科技"

markNameMap["libapssec.so"]= "盛大"

markNameMap["librsprotect.so"]= "瑞星"

def check_jiagu(filename):
    azip = zipfile.ZipFile(filename) # 默认模式r,读
    # print(azip.namelist()) # 返回所有文件夹和文件
    # print(azip.filename)# 返回该zip的文件名
    jiagu="未检测到加固"
    stop=False
    for path in markNameMap.keys():
        if stop:
            break
        else:
            for zippath in azip.namelist():
                if path==zippath:
                    jiagu=u"检测到 【{}】 加固".format(markNameMap[path])
                    stop = True
                    break
    print(jiagu)

if __name__ == '__main__':
    print(u'基于 "ApkScan-PKID查壳工具" 的规则修改 by 6time')
    print("""目前仅支持的加固厂商：
        娜迦、梆梆、爱加密、通付盾、360加固、
        百度加固、阿里加固、腾讯加固、 网易易盾
        盛大加固、瑞星加固、网秦加固、
        国信灵通加固、几维安全加固、顶像科技加固、
        apkprotect加固、等常规厂商
        """)
    print('python3 apkscan-pkid.py D:\\android\\app.apk')
    check_jiagu(sys.argv[1])
