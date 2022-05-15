# gRPC

https://grpc.io/docs/languages/go/quickstart/
https://github.com/grpc/grpc.io/blob/master/content/docs/languages/go/quickstart.md

export GOPATH=~/dev/languages/go/grpc
export GOROOT=/opt/homebrew/Cellar/go/1.18.2/libexec

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
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
go run src/greeter_client/main.go --name test

task run-server
task run-client -- --name test

## Update
### Regenerate gRPC code

../data/protoc/bin/protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/helloworld/helloworld.proto

## Run again