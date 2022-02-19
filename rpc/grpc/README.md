# golang 开发 grpc 文档

## 安装

### 工具及插件

* `protoc 下载`: <https://github.com/protocolbuffers/protobuf/releases>

* `proto 编译插件`:

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

* `grpc-gateway 插件`:

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

### 编译及使用

* 编译服务: just grpc => server/main.go

```bash
$cd rpc/grpc/protos/helloworld
$protoc --go_out=plugins=grpc:. helloworld.proto
```

* * 启动服务

  ```bash
  $cd rpc/grpc/server
  $go run main.go
  ```

* * 访问服务

  ```bash
  # grpc
  $cd rpc/grpc/client
  $go run main.go
  ```

* 编译服务: both grpc and http(gateway) => serverwithhttp/main.go

```bash
$cd rpc/grpc/
$protoc -I ./protos \
    --go_out ./protos --go_opt paths=source_relative \
    --go-grpc_out ./protos --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./protos --grpc-gateway_opt paths=source_relative \
    ./protos/helloworldwithhttp/helloworldwithhttp.proto
```

* * 启动服务

  ```bash
  $cd rpc/grpc/serverwithhttp
  $go run main.go
  ```

* * 访问服务

  ```bash
  # grpc
  $cd rpc/grpc/clientwithhttp
  $go run main.go

  # http
  $ curl -X POST -k http://localhost:8000/v1/greeter/sayhello -d '{"name": "world"}'
  ```
