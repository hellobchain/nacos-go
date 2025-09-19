VERSION=v1.0.0
build:
	@cd cmd && go mod tidy && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o ../bin/nacos-go.bin

build-linux:
	@cd cmd && go mod tidy && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o ../bin/nacos-go.bin

docker-build:
	@docker build -t github.com/hellobchain/nacos-go:${VERSION} -f ./docker/Dockerfile .

docker-build-linux:
	@cd cmd && go mod tidy && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o ../bin/nacos-go.bin
	@docker build -t github.com/hellobchain/nacos-go:${VERSION} -f ./docker/Dockerfile-linux .
