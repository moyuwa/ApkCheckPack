#!/bin/bash

echo "开始编译ApkCheckPack..."

# 创建bin目录（如果不存在）
mkdir -p bin

# 清理之前的编译结果
rm -f bin/ApkCheckPack_* 2>/dev/null

# 编译Windows amd64版本
echo "编译Windows amd64版本..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ApkCheckPack_windows_amd64.exe ./src

# 编译Linux amd64版本
echo "编译Linux amd64版本..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ApkCheckPack_linux_amd64 ./src

# 编译macOS amd64版本
echo "编译macOS amd64版本..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/ApkCheckPack_darwin_amd64 ./src

# 显示编译结果
echo "编译完成！"
echo "生成的文件："
ls -la bin/ApkCheckPack_*