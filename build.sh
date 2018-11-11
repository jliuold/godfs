#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. godfs.proto
cd ./server/main/
go build -o ../../godfs
