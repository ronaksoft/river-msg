#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./go/msg/*.pb.go
rm ./go/msg/*_helpers.go

# Create 'msg' package
cd ./proto/ || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gohelpers_out=../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gokk_out=../go/msg ./*.proto
cd "$currentWorkingDir"/go/msg || exit
go fmt
cd "$currentWorkingDir" || exit