.PHONY: build

# build
build: build-http

build-http:
	@echo " > Building [app]..."
	@cd ./cmd/http/ && go build && cd ../..
	@echo " > Finished building [app]"

# test
test:
	@echo " > Testing all..."
	@export SYSENV=testing && go test -race -cover ./...
	@echo " > Finished testing"

test-clear:
	@echo " > Testing all..."
	@export SYSENV=testing && go test -race -cover -gcflags="-l" ./...
	@echo " > Finished testing"

test-v:
	@echo " > Testing all; with verbose..."
	@export SYSENV=testing && go test -race -cover -v ./...
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
	@docker build -t dummy .
	@docker-compose up