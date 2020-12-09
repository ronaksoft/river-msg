#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./go/msg/*.pb.go
rm ./go/msg/*_helpers.go
rm ./go/rony/msg/*.pb.go
rm ./go/rony/msg/*.rony.go

# Create 'msg' package
cd ./proto/ || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gogofaster_out=plugins=grpc:../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gohelpers_out=../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  -I="${GOPATH}"/src --gokk_out=../go/msg ./*.proto
cd "$currentWorkingDir"/go/msg || exit
go fmt
cd "$currentWorkingDir"


## Create 'msg' package compatible with 'Rony'
## For Rony we must ignore lines
# 1. option (gogoproto.goproto_enum_prefix_all) = false;
# 2. import "github.com/gogo/protobuf/gogoproto/gogo.proto";
cd ./proto || exit
regex='^(option \(gogoproto\.goproto_enum_prefix_all\) = false;|import "github\.com/gogo/protobuf/gogoproto/gogo\.proto";)$'
mkdir tmp
for proto in ./*.proto; do
  grep -Ev "$regex" "$proto" >> ./tmp/"$proto"
done
cd ./tmp || exit
protoc  -I=.  --go_out=../../go/rony/msg ./*.proto
protoc  -I=.  --gorony_out=../../go/rony/msg ./*.proto
protoc  -I=.  --gokk_out=../../go/rony/msg ./*.proto
cd "$currentWorkingDir"/go/rony/msg || exit
go fmt
cd "$currentWorkingDir" || exit
rm -r "$currentWorkingDir"/proto/tmp
