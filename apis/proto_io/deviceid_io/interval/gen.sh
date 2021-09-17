#!/usr/bin/env bash

# 执行上面命令之后会在 $GOPATH/bin目录生成protoc-gen-go,protoc-gen-go-grpc两个文件
#go get google.golang.org/protobuf/cmd/protoc-gen-go
#go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --version

# grpc
function gen() {
  protoc -I . --go_out=paths=source_relative,plugins=grpc:. $1

  echo $1
}

gen deviceid.proto && mv deviceid.pb.go ../

