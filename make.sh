#!/usr/bin/env bash
protoc \
	-I=pkg/rpc/proto \
	-I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:$GOPATH/src \
	--grpc-gateway_out=logtostderr=true:$GOPATH/src \
	--govalidators_out=$GOPATH/src \
	health.proto

go build -o ./bin/rpc ./cmd/rpc/rpc.go
go build -o ./bin/web ./cmd/web/web.go
go build -o ./bin/worker ./cmd/worker/worker.go
