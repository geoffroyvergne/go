# Route_Guide GRPC

## Generate gRPC code
../data/protoc/bin/protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/routeguide/route_guide.proto

## Create module
go mod init example.com/go/grpc/routeguide

## Install dependencies
go get ./...

## Run

go run src/server/server.go

go run src/client/client.go

