#!/usr/bin/env bash
for protofile in $(find pkg/rpc/proto -name "*.proto"); do
	protoc \
		-I=pkg/rpc/proto \
		-I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:$GOPATH/src \
		--grpc-gateway_out=logtostderr=true:$GOPATH/src \
		--govalidators_out=$GOPATH/src \
		$protofile
done

go build -o ./bin/rpc ./cmd/rpc/rpc.go
go build -o ./bin/web ./cmd/web/web.go
go build -o ./bin/worker ./cmd/worker/worker.go
