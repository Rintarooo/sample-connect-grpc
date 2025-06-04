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
