.PHONY: protoc

all: protoc
protoc:
	protoc --go_out=plugins=grpc:./lib hello.proto

clean:
	rm -rf lib/*.go
