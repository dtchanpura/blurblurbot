# Blurblurbot
[![Build Status](https://travis-ci.org/dtchanpura/blurblurbot.svg?branch=master)](https://travis-ci.org/dtchanpura/blurblurbot)

Telegram bot for manipulating images accessible at [@blurblurbot](https://t.me/blurblurbot)

Tests yet to be added.

## Running

```bash
$ go get -u github.com/dtchanpura/blurblurbot
$ export TG_BOT_TOKEN=123456:abcdef1234
$ blurblurbot
2017/09/13 23:14:15 Listening on :8080
```

## Resources

This project uses mainly two third-party libraries.

* [esimov/stackblur-go](https://github.com/esimov/stackblur-go)
* [disintegration/imaging](https://github.com/disintegration/imaging)

Dockerfile has been created by referring

* [syncthing/docker](https://github.com/syncthing/docker)
* [chihaya/chihaya](https://github.com/chihaya/chihaya)
