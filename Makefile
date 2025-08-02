.PHONY: help run proto

help: ## Print help
	@awk 'BEGIN {FS=":.*##";printf"Makefile\n\nUsage:\n  make [command]\n\nAvailable Commands:\n"}/^[a-zA-Z_0-9-]+:.*?##/{printf"  %-40s%s\n",$$1,$$2}/^##@/{printf"\n%s\n",substr($$0,5)}' $(MAKEFILE_LIST)


run: ## Run the app locally
	@go run ./cmd/main.go

port-forward: ## Kubernetes port forward
	@kubectl port-forward service/api-stats 50051:50051 -n api-stats
