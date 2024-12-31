# Description: Makefile for labs observability and open telemetry

MS_INPUTVALIDATE_DIR=ms-inputvalidate
MS_WEATHERZIP_DIR=ms-weatherzip

.DEFAULT_GOAL := help

.PHONY: generate-docs-inputvalidate generate-docs-weatherzip generate-docs test-inputvalidate test-weatherzip tests

help: ## Exibe este menu de ajuda
	@echo "Opções disponíveis no Makefile:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

generate-docs-inputvalidate: ## Gera documentação Swagger para o serviço ms-inputvalidate
	@echo "Generating Swagger documentation for InputValidate..."
	@cd $(MS_INPUTVALIDATE_DIR) && swag init -g cmd/inputvalidate/main.go --output ./doc/swagger

generate-docs-weatherzip: ## Gera documentação Swagger para o serviço ms-weatherzip
	@echo "Generating Swagger documentation for WeatherZip..."
	@cd $(MS_WEATHERZIP_DIR) && swag init -g cmd/weatherzip/main.go --output ./doc/swagger

generate-docs: generate-docs-inputvalidate generate-docs-weatherzip ## Gera documentação Swagger para todos oso serviços

test-inputvalidate: ## Executa suite de testes do serviço ms-inputvalidate
	@echo "Running tests for InputValidate..."
	@cd $(MS_INPUTVALIDATE_DIR) && go test -timeout 30s -v ./... -coverprofile=report/coverage.out
	@cd $(MS_INPUTVALIDATE_DIR) && go tool cover -func=report/coverage.out

test-weatherzip: ## Executa suite de testes do serviço ms-weatherzip
	@echo "Running tests for WeatherZip..."
	@cd $(MS_WEATHERZIP_DIR) && go test -timeout 30s -v ./... -coverprofile=report/coverage.out
	@cd $(MS_WEATHERZIP_DIR) && go tool cover -func=report/coverage.out

tests: test-inputvalidate test-weatherzip ## Executa suite de testes de todos os serviços