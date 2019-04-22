#!/usr/bin/env bash
go build -o ./bin/rpc ./cmd/rpc/rpc.go
go build -o ./bin/web ./cmd/web/web.go
go build -o ./bin/worker ./cmd/worker/worker.go


