# Server-side gRPC Client for REST [Laravel]

### Generate protoc
```shell
protoc --proto_path=../go-grpc/config/grpc \
    --php_out=./config --plugin=/usr/local/bin/grpc_php_plugin \
    ../go-grpc/config/grpc/market.proto
```
