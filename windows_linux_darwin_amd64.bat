echo "windows_amd64"
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -o bin/ApkCheckPack_windows_amd64.exe ./src
echo "linux_amd64"
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags="-s -w" -o bin/ApkCheckPack_linux_amd64 ./src
echo "darwin_amd64"
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags="-s -w" -o bin/ApkCheckPack_darwin_amd64 ./src

rem Linux上跨平台编译程序
rem CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o your-program.exe your-program.go
rem CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o your-program.exe your-program.go
rem CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o your-program your-program.go

