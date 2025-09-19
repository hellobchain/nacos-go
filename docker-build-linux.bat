@echo off
setlocal
set VERSION=v1.0.0

:: 1) 交叉编译 Linux 二进制
cd /d "%~dp0cmd"
go mod tidy
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -trimpath -ldflags="-s -w" -o ..\bin\nacos-go.bin

:: 2) 构建镜像
cd /d "%~dp0"
docker build -t github.com/hellobchain/nacos-go:%VERSION% -f .\docker\Dockerfile-linux .