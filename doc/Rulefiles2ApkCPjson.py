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
    # print('extract_filename_jsoncontent')
    '''
    libMSDK.so.json

{
    "data": [
        {
            "locale": "zh-Hans",
            "data": {
                "label": "腾讯游戏 MSDK",
                "dev_team": "Tencent",
                "rule_contributors": [
                    "makazeu",
                    "crrashh1542",
                    "yfdyh000"
                ],
                "description": "MSDK，全称 Multi-platform Game Software Development Kit 是腾讯游戏公共组件和服务库平台，由增值服务部&公共数据平台部开发。为腾讯研发、发行的海内外游戏产品（移动、PC、主机平台）提供包括账号服务、好友关系链等必要增值服务能力接入。",
                "source_link": "https://docs.itop.qq.com/"
            }
        },
        {
            "locale": "en",
            "data": {
                "label": "Tencent Game MSDK",
                "dev_team": "Tencent",
                "rule_contributors": [
                    "makazeu",
                    "crrashh1542"
                ],
                "description": "MSDK, short for Multi-platform Game Software Development Kit, is Tencent's public component and service library platform developed by the Value-Added Services Department & Public Data Platform Department. It provides necessary value-added service capabilities such as account services and friend relationship chains for Tencent's R&D and published game products (mobile, PC, console platforms) at home and abroad.",
                "source_link": "https://docs.itop.qq.com/"
            }
        }
    ],
    "uuid": "51CC509E-7445-4D62-AAA4-36AC45C26A78"
}
    '''
    # 提取文件名 libMSDK.so.json -> libMSDK.so
    filename = os.path.basename(filepath)
    libstr = filename.replace('.json', '')
    # 提取JSON内容
    with open(filepath, 'r', encoding='utf-8') as f:
        jsoncontent = json.load(f)

    # 提取多语言数据，包括label、dev_team和description
    zh_data = {}
    en_data = {}
    
    for item in jsoncontent.get('data', []):
        locale = item.get('locale')
        data = item.get('data', {})
        if locale == 'zh-Hans':
            zh_data = {
                'label': data.get('label', ''),
                'dev_team': data.get('dev_team', ''),
                # 'description': data.get('description', '')
            }
        elif locale == 'en':
            en_data = {
                'label': data.get('label', ''),
                'dev_team': data.get('dev_team', ''),
                # 'description': data.get('description', '')
            }
    
    # 构建新的JSON结构，按中英文分开，移除UUID
    result = {
        'soname': libstr,
        'zh': {
            'label': zh_data.get('label', ''),
            'dev_team': zh_data.get('dev_team', ''),
            # 'description': zh_data.get('description', '')
        },
        'en': {
            'label': en_data.get('label', ''),
            'dev_team': en_data.get('dev_team', ''),
            # 'description': en_data.get('description', '')
        }
    }

    return result


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
