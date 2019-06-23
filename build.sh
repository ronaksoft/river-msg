#!/usr/bin/env bash

currentWorkingDir=`pwd`
rm ./ext/*.pb.go

# Create 'msg' package
cd ./ext/proto/
protoc  -I=. -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf -I=${GOPATH}/src --gogofaster_out=../ ./*.proto

# Create Constructors
cd ..
go run ../tools/constructor-builder/main.go . msg
go fmt

# Create 'internalMsg' package
cd ${currentWorkingDir}
rm ./int/*.pb.go
cd ./int/proto/
protoc  -I=${GOPATH}/src/git.ronaksoftware.com -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf -I=${GOPATH}/src -I=. --gogofaster_out=../ ./*.proto

# Create Constructors
cd ..
go run ../tools/constructor-builder/main.go . internalMsg
rm gen_results.go
go fmt

# Create Notification gRPC Stubs
cd ${currentWorkingDir}/notif/proto/
protoc -I=. -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf -I=${GOPATH}/src --gogofaster_out=plugins=grpc:../ ./*.proto
