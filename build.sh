#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./chat/*.pb.go

# Create 'msg' package
cd ./chat/proto/ || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gohelpers_out=../ ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gokk_out=../ ./*.proto
cd ..
go fmt


# Create Notification gRPC Stubs
cd "${currentWorkingDir}"/notif/proto/ || exit
protoc -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
