.PHONY: run proto

run:
	@go run ./cmd/main.go

proto:
	@export PATH="$PATH:/Users/ethantlee/go/bin"
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/stats.proto

port-forward:
	@kubectl port-forward service/api-stats 50051:50051 -n api-stats
