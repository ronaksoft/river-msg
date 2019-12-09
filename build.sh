#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./ext/*.pb.go

# Create 'msg' package
cd ./ext/proto/ || exit
protoc  -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
protoc  -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gohelpers_out=../ ./*.proto
protoc  -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gokk_out=../ ./*.proto
cd ..
go fmt


# Create 'admin' package
cd "${currentWorkingDir}" || exit
rm ./admin/*.pb.go
cd ./admin/proto/ || exit
protoc  -I="${currentWorkingDir}"/../vendor -I="${GOPATH}"/src/git.ronaksoftware.com -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src -I=. --gogofaster_out=../ ./*.proto


# Create Notification gRPC Stubs
cd "${currentWorkingDir}"/notif/proto/ || exit
protoc -I="${currentWorkingDir}"/../vendor -I=. -I="${GOPATH}"/src/github.com/gogo/protobuf/protobuf -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
