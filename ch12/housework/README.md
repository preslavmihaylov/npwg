# Instructions

## Generate protobuf
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative idl/housework.proto
```

## Build
```
go build -o hw_server *.go
go build -o hw_client client/*.go
```

## Run standalone
```
./hw_server --format {json, gob, protobuf} {complete <item-id>, add <desc>, list}
```

## Run server & client
```
./hw_server serve
./hw_client {complete <item-id>, add <desc>, list}
```
