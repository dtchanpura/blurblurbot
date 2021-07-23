BUILD_DATE := $(shell date)

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)


SED_CMD = sed -i.bak

UNAME_S := $(shell uname -s)

BINARY := blurblurbot

PACKAGES := $(shell go list ./... | grep -v /vendor)

all: build silent-test

build:
	go build -ldflags "-s -w -X github.com/dtchanpura/blurblurbot/bot.Version=`git tag --points-at=HEAD` -X github.com/dtchanpura/blurblurbot/bot.BuildDate=`date --rfc-3339 date`" -o bin/$(BINARY)-$(GOOS)-$(GOARCH)

test:
	TG_BOT_TOKEN="dummy" go test -v $(PACKAGES)

silent-test:
	TG_BOT_TOKEN="dummy" go test $(PACKAGES)

format:
	go fmt $(PACKAGES)
