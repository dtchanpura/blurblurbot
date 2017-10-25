BUILD_DATE := $(shell date)

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)


SED_CMD = sed -i

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
        SED_CMD += ""
endif

BINARY := blurblurbot

PACKAGES := $(shell go list ./... | grep -v /vendor)

DEPENDENCIES := \
	github.com/disintegration/imaging \
	github.com/esimov/stackblur-go


all: build silent-test

build: version deps
	go build -o bin/$(BINARY)-$(GOOS)-$(GOARCH) main.go

version:
	@$(SED_CMD) 's/BuildDate = .*/BuildDate = "$(BUILD_DATE)"/' bot/version.go

test:
	TG_BOT_TOKEN="dummy" go test -v $(PACKAGES)

silent-test:
	TG_BOT_TOKEN="dummy" go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
