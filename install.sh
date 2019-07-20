#!/usr/bin/env bash

############
### Make sure you have protoc installed on your system.
##### 2

go get -u github.com/gogo/protobuf/proto
go get -u github.com/gogo/protobuf/jsonpb
go get -u github.com/gogo/protobuf/protoc-gen-gogo
go get -u github.com/gogo/protobuf/protoc-gen-gogofast
go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
go get -u github.com/gogo/protobuf/gogoproto