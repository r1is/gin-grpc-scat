#!/bin/bash
export GOPROXY=https://goproxy.io,direct
export CGO_ENABLED=0

cd analyze_tool
go mod tidy
go build -o ./build/main ./main.go

cd ../contract_analyzer
go mod tidy
go build -o ./build/main ./main.go