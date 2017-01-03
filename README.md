# Compress Middleware

[![Build Status](https://travis-ci.org/go-gem/middleware-compress.svg?branch=master)](https://travis-ci.org/go-gem/middleware-compress)
[![GoDoc](https://godoc.org/github.com/go-gem/middleware-compress?status.svg)](https://godoc.org/github.com/go-gem/middleware-compress)
[![Coverage Status](https://coveralls.io/repos/github/go-gem/middleware-compress/badge.svg?branch=master)](https://coveralls.io/github/go-gem/middleware-compress?branch=master)

Compress middleware for [Gem](https://github.com/go-gem/gem) Web framework.

## Getting Started

**Install**

```
$ go get -u github.com/go-gem/middleware-compress
```

**Compress Levels**

- gzip.DefaultCompression
- gzip.BestSpeed
- gzip.BestCompression

Note: the other levels would trigger panic.

**Example**

```
package main

import (
	"compress/gzip"
	"log"
	"net/http"
	"os"

	"github.com/go-gem/gem"
	"github.com/go-gem/middleware-compress"
)

func main() {
	compressMidware := compressmidware.New(gzip.BestCompression)

	router := gem.NewRouter()
	router.ServeFiles(
		"/tmp1/*filepath", http.Dir(os.TempDir()),
		&gem.HandlerOption{Middlewares: []gem.Middleware{compressMidware}},
	)
	router.ServeFiles("/tmp2/*filepath", http.Dir(os.TempDir()))

	log.Println(gem.ListenAndServe(":8080", router.Handler()))
}
```