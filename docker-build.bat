@echo off
setlocal
set VERSION=v1.0.0
docker build -t github.com/hellobchain/nacos-go-web:%VERSION% -f .\docker\Dockerfile-front .
docker build -t github.com/hellobchain/nacos-go:%VERSION% -f .\docker\Dockerfile .