#!/bin/bash

echo "Build app"

[ -f $GOBIN/testapp ] && rm $GOBIN/testapp

cd $GOPATH && go install ./src/testapp/testlib/
go build -o $GOBIN/testapp ./src/testapp

$GOBIN/testapp
