# gRPC Server [Golang]

### Generate protoc
```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    config/grpc/market.proto
```