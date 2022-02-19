# golang 开发 grpc 文档

## grpc 目录

```plain
├── README.md
├── client
│   └── main.go
├── clientwithhttp
│   └── main.go
├── protos
│   ├── google
│   │   └── api
│   ├── helloworld
│   │   ├── helloworld.pb.go
│   │   └── helloworld.proto
│   └── helloworldwithhttp
│       ├── helloworldwithhttp.pb.go
│       ├── helloworldwithhttp.pb.gw.go
│       ├── helloworldwithhttp.proto
│       └── helloworldwithhttp_grpc.pb.go
├── server
│   └── main.go
└── serverwithhttp
    └── main.go
```

## 安装

### 工具及插件

* `protoc 下载`: <https://github.com/protocolbuffers/protobuf/releases>

* `proto 编译插件`:

```bash
# protoc 编译 go
$go get -u github.com/golang/protobuf/protoc-gen-go
$go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
# grpc-gateway: grpc -> http
$go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

* 下载 `grpc-gateway` 编译用的依赖文件(存放于 `protos` 目录)

  * 文件一: google/api/annotations.proto

    * 下载路径: <https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto>
  
  * 文件二: google/api/http.proto

    * 下载路径: <https://github.com/googleapis/googleapis/blob/master/google/api/http.proto>

### 编译及使用

注意: 下面所有 `cd` 之前的目录位于 `grpc` 目录

* just grpc => server/main.go

```bash
# 编译
$cd protos/helloworld
$protoc --go_out=plugins=grpc:. helloworld.proto

# 启动 grpc 服务
$cd server
$go run main.go

# 访问 grpc 服务
$cd client
$go run main.go
```

* both grpc and http(gateway) => serverwithhttp/main.go

```bash
# 编译
$cd .
$protoc -I ./protos \
    --go_out ./protos --go_opt paths=source_relative \
    --go-grpc_out ./protos --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./protos --grpc-gateway_opt paths=source_relative \
    ./protos/helloworldwithhttp/helloworldwithhttp.proto

# 启动 grpc & http 服务
$cd serverwithhttp
$go run main.go

# 访问 grpc 服务
$cd clientwithhttp
$go run main.go

# 访问 http 服务
$curl -X POST -k http://localhost:8000/v1/greeter/sayhello -d '{"name": "world"}'
```
