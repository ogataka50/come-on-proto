.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate-all
generate-all: ## generate code from proto
	$(MAKE) generate-go
	$(MAKE) generate-php

.PHONY: generate-go
generate-go: ## generate go code from proto
	protoc --go_out=plugins=grpc:./go protos/*.proto

.PHONY: generate-php
generate-php: ## generate php code from proto
	protoc --proto_path=protos/ \
		--php_out=php \
		--grpc_out=php \
		--plugin=protoc-gen-grpc=./grpc_php_plugin \
		./protos/*.proto