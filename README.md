***Hide***, *what You don't want the audience to* ***see*** - **[Daryl](Daryl.md)**

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangsam/internal)](https://goreportcard.com/report/github.com/GoLangsam/internal)
[![Build Status](https://travis-ci.org/GoLangsam/internal.svg?branch=master)](https://travis-ci.org/GoLangsam/internal)
[![GoDoc](https://godoc.org/github.com/GoLangsam/internal?status.svg)](https://godoc.org/github.com/GoLangsam/internal)

In *[go](http://golang.org)* any package below/inside `internal` is **not** exported/published to any package above - (only to siblings and their descendants).

Thus: A great place to 'hide' stuff in plain sight :-)

Stuff that is either
- not good enough (yet) or
- not general enough (yet)
in order to be used in Your project.

---
Stuff such as:

## cmd/
- `dotpath` - simple CLI to play with `./container/ccsafe/dotpath`
- `glob` - simple CLI to play with `filepath/glob`

---
## container/

### ccsafe/
- `dotpath` - a fs-plist with more meaning of dots

### oneway
no packages yet

---
## do/
- `dot`

### do/cmd/
- `cancel` - a brute-force Cancellor for CLI `cmd`s

---
Your suggestions, remarks, questions and/or contributions are welcome ;-)

---
## Think deep - code happy - be simple - see clear :-)

---
## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/internal/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/internal)  [![Beerpay](https://beerpay.io/GoLangsam/internal/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/internal?focus=wish)