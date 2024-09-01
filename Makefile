.PHONY: build

# build
build: build-http

build-http:
	@echo " > Building [app]..."
	@cd ./cmd/http/ && go build && cd ../..
	@echo " > Finished building [app]"

# test

test-coverage:
	@echo " > Testing all; with verbose..."
	@export SYSENV=testing && go test -race -cover -v ./... -coverprofile=cover.out && ./exclude-code-coverage.sh
	@echo "> ******* coverage of function *******"
	@go tool cover -func cover.out
	@echo " > Finished testing"

test-fail:
	@echo " > Testing all; with verbose; show fail only..."
	@export SYSENV=testing && go test -race -cover -v ./... | grep -A10 -B2 'FAIL\:'
	@echo " > Finished testing"

# run
run-http: build-http
	@echo " > Running [app]..."
	@cd ./cmd/http/ && ./http
	@echo " > Finished running [app]"

run-docker:
	@echo "> Build image on docker..."
	@docker build -t reconciliation-service .
	@docker-compose up