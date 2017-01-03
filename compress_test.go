// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package compressmidware

import (
	"compress/gzip"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-gem/gem"
)

func TestCompress(t *testing.T) {
	compressMidware := New(gzip.BestCompression)
	handler := compressMidware.Wrap(gem.HandlerFunc(func(ctx *gem.Context) {}))

	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	ctx := &gem.Context{Response: resp, Request: req}
	handler.Handle(ctx)

	if resp.Header().Get("Vary") != "Accept-Encoding" {
		t.Errorf("expected response header Vary: %q, got %q", "Accept-Encoding", resp.Header().Get("Vary"))
	}
}

func TestNew(t *testing.T) {
	// invalid compress level.
	level := gzip.BestCompression + 1
	expectErr := fmt.Errorf("invalid compression level requested: %d", level)

	defer func() {
		v := recover()
		if v != nil && reflect.DeepEqual(v, expectErr) {
			return
		}
		t.Errorf("expected err %q, got %q", expectErr, v)
	}()

	_ = New(gzip.BestCompression + 1)
}
