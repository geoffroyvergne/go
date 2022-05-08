# gRPC

https://grpc.io/docs/languages/go/quickstart/
https://github.com/grpc/grpc.io/blob/master/content/docs/languages/go/quickstart.md

export GOPATH=~/dev/languages/go/grpc
export GOROOT=/opt/homebrew/Cellar/go/1.18.1/libexec

go get -u github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"

## Install protoc
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-osx-x86_64.zip

export PATH="$PATH:$(go env GOPATH)/bin"

## Check port use
sudo lsof -i -n -P

## Run

go mod init example.com/go/grpc/helloworld

go run src/greeter_client/main.go

go get ./...

go run greeter_server/main.go
go run greeter_client/main.go

## Update
### Regenerate gRPC code

../data/protoc/bin/protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative helloworld/helloworld.proto

## Run again