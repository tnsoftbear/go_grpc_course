.PHONY: all
all: get-protoc-plugins proto-gen

.PHONY: get-protoc-plugins
get-protoc-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/streaming.proto

.PHONY: run-server
run-server:
	go run -mod=readonly cmd/server/main.go

.PHONY: run-client
run-client:
	go run -mod=readonly cmd/client-stream-client/main.go

run-server-stream-client:
	go run -mod=readonly cmd/server-stream-client/main.go

run-bi-client:
	go run -mod=readonly cmd/bi-directional-stream-client/main.go