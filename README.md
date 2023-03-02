# Go support for Protocol Buffers

[![Go Reference](https://pkg.go.dev/badge/google.golang.org/protobuf.svg)](https://pkg.go.dev/google.golang.org/protobuf)
[![Build Status](https://travis-ci.org/protocolbuffers/protobuf-go.svg?branch=master)](https://travis-ci.org/protocolbuffers/protobuf-go)

This project hosts the Go implementation for
[protocol buffers](https://protobuf.dev)

- [README.md](./README-EN.md)

- 2023.03.02 `protoc-gen-go` 新增参数  json_required=true 去掉结构体中 json tag `omitempty` 选项 

eg: protoc --proto_path=. --go_out=paths=source_relative,json_required=true:./ ./*.proto