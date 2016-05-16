# ---------------------------------------------------------------------------
# Makefile for GO utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=`git describe --tags 2>/dev/null`
LDFLAGS=-ldflags "-X main.version=${BUILD_TAG} -s -w"

init:
	sed -i "s/PROJECT_NAME/${PROJECT_DIR}/g" CHANGELOG.md README.md main.go

build: get
	go build ${LDFLAGS}

get:
	go get

test: fmt vet
	go test -v -cover

fmt:
	go fmt

vet:
	go vet -v

install:
	go install -a ${LDFLAGS} ./...


dist: clean build
	upx -9 ${PROJECT_DIR}.exe

clean:
	go clean
