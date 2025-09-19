# nacos-go

nacos-go 是一个基于 Go 语言开发的 Nacos 服务。

## 内存启动

### 1. 无需数据库，直接启动

go run . -driver memory

### 2. 发布配置

curl -X POST 127.0.0.1:8848/v1/cs/configs \
     -d '{"dataId":"app.yaml","group":"DEFAULT_GROUP","content":"user:\n  name: nacos"}'

### 3. 获取配置

curl "127.0.0.1:8848/v1/cs/configs?dataId=app.yaml&group=DEFAULT_GROUP"

### 4. 删除配置

curl -X DELETE "127.0.0.1:8848/v1/cs/configs?dataId=app.yaml&group=DEFAULT_GROUP"
