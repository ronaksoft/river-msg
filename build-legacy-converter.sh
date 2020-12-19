#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./go/msg/*.pb.go
rm ./go/msg/*.rony.go
rm ./go/msg/*.legacy.go

## Create 'msg' package compatible with 'Rony'
## For Rony we must ignore lines
cd ./proto || exit
regex='^(option \(gogoproto\.goproto_enum_prefix_all\) = false;|import "github\.com/gogo/protobuf/gogoproto/gogo\.proto";)$'
mkdir tmp
for proto in ./*.proto; do
  grep -Ev "$regex" "$proto" >> ./tmp/"$proto"
done
cd ./tmp || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  --golegacy_out=paths=source_relative:../../go/_legacy ./*.proto

cd "$currentWorkingDir"/go/msg || exit
go fmt
cd "$currentWorkingDir" || exit
rm -r "$currentWorkingDir"/proto/tmp
rm -r "$currentWorkingDir"/go/msg/github.com