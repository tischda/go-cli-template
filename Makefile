# ---------------------------------------------------------------------------
# Makefile for go CLI utilities
# ---------------------------------------------------------------------------

REPOSITORY=github.com/tischda
PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=$(shell git describe --tags)
LDFLAGS=-ldflags=all="-X main.version=${BUILD_TAG} -s -w"

all: build

init:
	sed -i "s/PROJECT_NAME/${PROJECT_DIR}/g" CHANGELOG.md README.md appveyor.yml .travis.yml main.go
	go mod init ${REPOSITORY}/${PROJECT_DIR}
	go mod tidy
	go mod vendor

build:
	go build ${LDFLAGS}

test:
	go test -v -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

install:
	go install ${LDFLAGS} ./...

dist: clean build
	upx -9 ${PROJECT_DIR}.exe

clean:
	go clean
