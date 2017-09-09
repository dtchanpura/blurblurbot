BINARY := blurblurbot

PACKAGES := $(shell go list ./... | grep -v /vendor)

DEPENDENCIES := \
	github.com/disintegration/imaging \
	github.com/esimov/stackblur-go


all: build silent-test

build:
	go build -o bin/$(BINARY)-$(GOOS)-$(GOARCH) main.go

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
