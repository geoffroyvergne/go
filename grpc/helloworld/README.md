# Helloworld gRPC

https://grpc.io/docs/languages/go/quickstart/
https://github.com/grpc/grpc.io/blob/master/content/docs/languages/go/quickstart.md

## Generate gRPC code
../data/protoc/bin/protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/helloworld/helloworld.proto

## Create module
go mod init example.com/go/grpc/helloworld

## Install dependencies
go get ./...

## Run
go run greeter_server/main.go

go run greeter_client/main.go
go run src/greeter_client/main.go --name test

task run-server
task run-client -- --name test


