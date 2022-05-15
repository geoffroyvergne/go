# GOlang GRPC

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

