.PHONY: help run proto

help: ## Print help
	@awk 'BEGIN {FS=":.*##";printf"Makefile\n\nUsage:\n  make [command]\n\nAvailable Commands:\n"}/^[a-zA-Z_0-9-]+:.*?##/{printf"  %-40s%s\n",$$1,$$2}/^##@/{printf"\n%s\n",substr($$0,5)}' $(MAKEFILE_LIST)


run: ## Run the app locally
	@go run ./cmd/main.go

proto: ## Generate protobuf files
	@export PATH="$PATH:/Users/ethantlee/go/bin"
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/stats.proto

port-forward: ## Kubernetes port forward
	@kubectl port-forward service/api-stats 50051:50051 -n api-stats
