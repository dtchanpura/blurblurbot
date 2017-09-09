BINARY := blurblurbot

PACKAGES := \
	github.com/dtchanpura/blurblurbot \
	github.com/dtchanpura/blurblurbot/bot \
	github.com/dtchanpura/blurblurbot/cmd
DEPENDENCIES := \
	github.com/disintegration/imaging \
	github.com/esimov/stackblur-go


all: build silent-test

build:
	go build -o bin/$(BINARY) main.go

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
