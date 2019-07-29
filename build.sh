#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./ext/*.pb.go

# Create 'msg' package
cd ./ext/proto/
protoc  -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gogofaster_out=../ ./*.proto


# Create Constructors
cd ..
go run ../tools/constructor-builder/main.go . msg
go fmt

# Create 'internalMsg' package
cd "${currentWorkingDir}" || exit
rm ./int/*.pb.go
cd ./int/proto/ || exit
protoc  -I="${currentWorkingDir}"/../vendor -I="${GOPATH}"/src/git.ronaksoftware.com -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src -I=. --gogofaster_out=../ ./*.proto

# Create Constructors
cd ..
go run ../tools/constructor-builder/main.go . internalMsg
rm gen_results.go
go fmt

# Create 'admin' package
cd "${currentWorkingDir}" || exit
rm ./admin/*.pb.go
cd ./admin/proto/ || exit
protoc  -I="${currentWorkingDir}"/../vendor -I="${GOPATH}"/src/git.ronaksoftware.com -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src -I=. --gogofaster_out=../ ./*.proto


# Create Notification gRPC Stubs
cd "${currentWorkingDir}"/notif/proto/ || exit
protoc -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
