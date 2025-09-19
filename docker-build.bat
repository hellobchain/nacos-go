@echo off
setlocal
set VERSION=v1.0.0
docker build -t dm/nacos-go:%VERSION% -f .\docker\Dockerfile .