#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./chat/*.pb.go
rm ./msg/*.pb.go
rm ./msg/*_helpers.go


# Create 'msg' package
cd ./proto/ || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gohelpers_out=../msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gokk_out=../msg ./*.proto
cd ../msg
go fmt

# Create Notification gRPC Stubs
cd "${currentWorkingDir}"/notif/proto/ || exit
protoc -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../ ./*.proto
