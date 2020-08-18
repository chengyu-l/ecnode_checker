#!/bin/bash

basepath=$(cd `dirname $0`; pwd)
export GOPATH=$GOPATH:$basepath
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build cmd/main.go