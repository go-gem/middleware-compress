// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package compressmidware is a HTTP middleware that compress response body.

Compress Levels

	gzip.DefaultCompression
	gzip.BestSpeed
	gzip.BestCompression

Example

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
*/
package compressmidware

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/go-gem/gem"
)

// New returns a Compress instance by
// the given compress level.
//
// Support levels:
// 	gzip.DefaultCompression
// 	gzip.BestSpeed
// 	gzip.BestCompression
//
// an invalid level would trigger panic.
func New(level int) *Compress {
	wrapper, err := gziphandler.NewGzipLevelHandler(level)
	if err != nil {
		panic(err)
	}
	return &Compress{wrapper: wrapper}
}

// Compress is a HTTP middleware that compress the response
// body.
type Compress struct {
	wrapper func(http.Handler) http.Handler
}

// Wrap implements the Middleware interface.
func (c *Compress) Wrap(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		c.wrapper(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.Response = w
			next.Handle(ctx)
		})).ServeHTTP(ctx.Response, ctx.Request)
	})
}
