.PHONY: proto
proto: pkg/perftest/perftest.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/perftest/perftest.proto

.PHONY: run-server
run-server: cmd/server/*
	go run cmd/server/*

.PHONY: run-client
run-client: cmd/server/*
	go run cmd/client/*
