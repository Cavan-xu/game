SHELL = /bin/bash

windows:
	protoc  --plugin=protoc-gen-go.exe   --go_out=../rpc game.proto common.proto

linux:
	protoc --go_out=../ common.proto game.proto
