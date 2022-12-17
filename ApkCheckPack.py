#!/usr/bin/env python3
# coding=utf-8
# python version 3.7 by 6time
#

import zipfile, hashlib, json
from tkinter import *
from tkinter.filedialog import askdirectory, askopenfilename
from tkinter.messagebox import showinfo
import windnd

# packs = """"""
# # 字符串方式加载特征
# markNameMap = json.loads(packs)
# markNameMap = dict(markNameMap)
# json文件方式加载特征
with open('apkpack.json', 'r', encoding='utf-8') as f:
    markNameMap = json.load(f)
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
    help = ("""ApkCheckPack v1.2 by 6time\n
    目前规则总数：177条
        """)
    root = Tk()
    root.title("ApkCheckPack v1.2 by 6time")
    # 获取屏幕 宽、高
    ws = root.winfo_screenwidth()
    hs = root.winfo_screenheight()
    root.geometry('%dx%d+%d+%d' % (600, 500, ws / 2 - 300, 0))
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
