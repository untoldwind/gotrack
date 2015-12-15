DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./server/...)
PACKAGES = $(shell go list ./server/...)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
         -nilfunc -printf -rangeloops -shift -structtags -unsafeptr
DEPPATH = $(firstword $(subst :, , $(GOPATH)))
VERSION = $(shell date -u +.%Y%m%d.%H%M%S)

all: export GOPATH=${PWD}/Godeps/_workspace:${PWD}/../../../..
all: format
	@mkdir -p bin/
	@echo "--> Running go build"
	@go build -ldflags "-X github.com/untoldwind/gotrack/server/config.versionMinor=${VERSION}" -v -o bin/gotrack github.com/untoldwind/gotrack/server

deps:
	@echo "--> Installing build dependencies"
	@go get -d -v ./server/... $(DEPS)

updatedeps: deps
	@echo "--> Updating build dependencies"
	@go get -d -f -u ./server/... $(DEPS)

test: export GOPATH=${PWD}/Godeps/_workspace:${PWD}/../../../..
test: deps
	@echo "--> Running tests"
	@go test -v ./server/...
	@$(MAKE) vet

format: export GOPATH=${PWD}/Godeps/_workspace:${PWD}/../../../..
format: deps
	@echo "--> Running go fmt"
	@go fmt ./server/...

raspberry: export GOPATH=${PWD}/Godeps/_workspace:${PWD}/../../../..
raspberry: export GOOS=linux
raspberry: export GOARCH=arm
raspberry: export CGO_ENABLED=0
raspberry:
	@mkdir -p bin/arm
	@echo "--> Running go build (linux, arm)"
	@go build -ldflags "-X github.com/untoldwind/gotrack/server/config.versionMinor=${VERSION}" -v -o bin/arm/gotrack github.com/untoldwind/gotrack/server

vet: export GOPATH=${PWD}/Godeps/_workspace:${PWD}/../../../..
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "--> Running go tool vet $(VETARGS)"
	@find server -name "*.go" | grep -v "./Godeps/" | xargs go tool vet $(VETARGS); if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for reviewal."; \
	fi

godepssave:
	@echo "--> Godeps save"
	@go build -v -o bin/godep github.com/tools/godep
	@bin/godep save ./server/...
