#!/usr/bin/env bash

currentWorkingDir=$(pwd)
rm ./go/msg/*.pb.go
rm ./go/msg/*.rony.go

## Create 'msg' package compatible with 'Rony'
## For Rony we must ignore lines
cd ./proto || exit
regex='^(option \(gogoproto\.goproto_enum_prefix_all\) = false;|import "github\.com/gogo/protobuf/gogoproto/gogo\.proto";)$'
mkdir tmp
for proto in ./*.proto; do
  grep -Ev "$regex" "$proto" >> ./tmp/"$proto"
done
cd ./tmp || exit
protoc  -I="${currentWorkingDir}"/vendor -I=.  --go_out=paths=source_relative:../../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  --gorony_out=paths=source_relative:../../go/msg ./*.proto
protoc  -I="${currentWorkingDir}"/vendor -I=.  --gokk_out=paths=source_relative:../../go/msg ./*.proto

cd "$currentWorkingDir"/go/msg || exit
go fmt
cd "$currentWorkingDir" || exit
rm -r "$currentWorkingDir"/proto/tmp
#rm -r "$currentWorkingDir"/go/_legacy/github.com