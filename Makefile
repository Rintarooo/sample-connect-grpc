.PHONY: help	
help:	
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'	

.PHONY: buf-format	
buf-format: ## Format proto files	
	docker compose run --rm buf-format	

.PHONY: buf-generate	
buf-generate: ## Generate source code from proto files	
	rm -rf backend/internal/generated/grpc
	rm -rf frontend/src/generated/ts/grpc/
	docker compose run --rm buf-generate

.PHONY: run-frontend
run-frontend: ## Run frontend locally
	cd frontend && yarn dev

.PHONY: run-backend
run-backend: ## Run backend locally
	cd backend && go run cmd/serve/main.go

.PHONY: go-mod-tidy
go-mod-tidy: ## Run go mod tidy
	cd backend && go mod tidy

.PHONY: curl-greet
curl-greet: ## Curl greet (usage: make curl-greet NAME=XXX)
	curl -X POST http://localhost:8080/greet.v1.GreetingService/GetGreeting \
		-H "Content-Type: application/json" \
		-d '{"name": "$(or $(NAME),YourName)"}'

.PHONY: grpcurl-greet
grpcurl-greet: ## Grpcurl greet (usage: make grpcurl-greet NAME=XXX)
	grpcurl -plaintext -d '{"name": "$(or $(NAME),YourName)"}' \
	  localhost:8080 greet.v1.GreetingService/GetGreeting
