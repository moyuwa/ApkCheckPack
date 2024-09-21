#!/usr/bin/env python3
# coding=utf-8
# python version 3.10
# SDK框架规则转换工具 by 6time
# Apk引用库检测对应规则库 https://github.com/LibChecker/LibChecker-Rules
# so库方式 LibChecker-Rules-master\native-libs

import json
import os, sys


# 枚举native-libs文件夹
def enum_native_libs(dirpath):
    print('enum_native_libs')
    file_list = []
    for root, dirs, files in os.walk(dirpath):
        for file in files:
            if file.endswith(".json") and ".so" in file:
                # 获取文件路径及名称
                file_path = os.path.join(root, file)  # 读取json内容需要路径
                # print(file_path)
                file_list.append(file_path)
    return file_list


# 提取文件名称和内容
def extract_filename_jsoncontent(filepath):
    print('extract_filename_jsoncontent')

    '''
    libMSDK.so.json

    {
        "label": "MSDK / MobileGame Software Development Kit",
        "team": "腾讯互动娱乐",
        "iconUrl": "",
        "contributors": ["yfdyh000"],
        "description": "MSDK，全称 MobileGame Software Development Kit，是增值服务部为游戏（主要为手游）提供公共组件和服务库的平台，旨在加快游戏接入各公共组件和海外开放平台。",
        "relativeUrl": "https://docs.msdk.qq.com/"
    }
    '''

    # 提取文件名 libMSDK.so.json -> libMSDK.so
    filename = os.path.basename(filepath)
    libstr = filename.replace('.json', '')
    # 提取JSON内容
    with open(filepath, 'r+', encoding='utf-8') as f:
        jsoncontent = json.load(f)

    # 提取所需字段
    # label = jsoncontent.get('label', '')
    # team = jsoncontent.get('team', '')
    # iconUrl = jsoncontent.get('iconUrl', '')
    # contributors = jsoncontent.get('contributors', [])
    # description = jsoncontent.get('description', '')
    # relativeUrl = jsoncontent.get('relativeUrl', '')

    # 加入额外字段
    jsoncontent['soname'] = libstr

    return jsoncontent


# 合并后写入json文件，规则转为ApkCheckPack可用数据
def com_save_json_apkcp():
    print('com_save_json_apkcp')
    # 将规则库内的文件夹放到同级目录
    dirpath = 'native-libs'
    file_list = enum_native_libs(dirpath)

    merged_json_content = []
    # 遍历文件列表
    for filepath in file_list:
        print('filepath:' + filepath)
        jsoncontent = extract_filename_jsoncontent(filepath)
        merged_json_content.append(jsoncontent)

    # 合并后写入json文件
    with open("merged_json_file.json", "w", encoding="utf-8") as merged_file:
        json.dump(merged_json_content, merged_file, indent=4, ensure_ascii=False)  # ensure_ascii=False 解决中文字符串写入问题

if __name__ == "__main__":
    print('convert-LibChecker-Rules')
    com_save_json_apkcp()
