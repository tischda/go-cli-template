# ---------------------------------------------------------------------------
# Makefile for go CLI utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=$(shell git describe --tags 2>/dev/null || echo undefined)
LDFLAGS=-ldflags=all="-X main.version=${BUILD_TAG} -s -w"

all: build
#--start-init--#
init:
	sed -i "s/##PROJECT_NAME##/${PROJECT_DIR}/g" CHANGELOG.md appveyor.yml .travis.yml main.go
	sed -i "s|##REPOSITORY##|${REPOSITORY}|g" README.tpl
	sed -i "s/##PROJECT_NAME##/${PROJECT_DIR}/g" README.tpl
	mv README.tpl README.md
	go mod init ${REPOSITORY}/${PROJECT_DIR}
	go mod tidy
	go mod vendor
	git init
#--end-init--#

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
