# gRPC

https://grpc.io/docs/languages/go/quickstart/

https://github.com/grpc/grpc.io/blob/master/content/docs/languages/go/quickstart.md

export GO111MODULE=on 
export GOPATH=$HOME/dev/go
export GOROOT=/usr/local/opt/go/libexec

go get github.com/golang/protobuf/protoc-gen-go@v1.3
export PATH="$PATH:$(go env GOPATH)/bin"

## Install protoc
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-osx-x86_64.zip

## Check port use
sudo lsof -i -n -P

## Run

go get ./...

go run greeter_server/main.go
go run greeter_client/main.go

## Update
### Regenerate gRPC code

../data/protoc/bin/protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative helloworld/helloworld.proto

## Run again