PKG_VERSION?=0.0.0
COMMIT=`git rev-parse --short HEAD`
PKG_LIST := $(shell go list ./... | grep -v -e /vendor/ -e /api/design -e /scan/logos -e /grpcapi/config/restore_validate )
GO_FILES := $(shell find . -name '*.go' | grep -v -e /vendor/ -e _test.goy)
SERVICE_NAME=sn

# docker build variables
CONTAINER_NAME=
VERSION_NAME=0.1

# flags
TEST_FLAGS=-count=1
BUILD_FLAGS=

.PHONY: lint, tests, coverage, build, docker, fmt, wire, api-generate, run, init

all:

lint: ## Lint the files
	@golangci-lint --timeout=2m run

tests:
	go test ${TEST_FLAGS} ./...

coverage:
	go test ${TEST_FLAGS} -coverprofile=coverage.out ./...
	go tool cover -html="coverage.out"

build: .
	go build ${BUILD_FLAGS} -o ${SERVICE_NAME}.exe cmd/main.go

docker: .
	docker build -f zarf/docker/Dockerfile -t ${CONTAINER_NAME}:${VERSION_NAME} .

migrate:
	go run cmd/main.go migrate

fmt:
	go fmt ./...

api-generate:
	go run -v -mod=mod  github.com/ogen-go/ogen/cmd/ogen -clean -package generated -config ./config/ogen_config.yml -target ./api/rest/generated ../config/openapi.json

wire:
	go run -v -mod=mod github.com/google/wire/cmd/wire ./api/inject

run:
	go run cmd/main.go api

init:
	find . -type f \( ! -regex '.*/\..*' ! -regex './Makefile' \) -exec sed -i "s/service_tpl_go/$(SERVICE_NAME)/g" {} +