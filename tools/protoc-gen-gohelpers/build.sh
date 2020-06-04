#!/usr/bin/env bash

protoc  -I=./testdata  --gohelpers_out=./testdata ./testdata/*.proto