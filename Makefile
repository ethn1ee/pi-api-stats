.PHONY: proto a

proto:
	@export PATH="$PATH:/Users/ethantlee/go/bin"
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/stats.proto
