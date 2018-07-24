// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cmd_go_bootstrap

package modfetch

import (
	"io"

	web "cmd/go/internal/web2"
)

// webGetGoGet fetches a go-get=1 URL and returns the body in *body.
// It allows non-200 responses, as usual for these URLs.
func webGetGoGet(url string, body *io.ReadCloser) error {
	return web.Get(url, web.Non200OK(), web.Body(body))
}

// webGetBytes returns the body returned by an HTTP GET, as a []byte.
// It insists on a 200 response.
func webGetBytes(url string, body *[]byte) error {
	return web.Get(url, web.ReadAllBody(body))
}

// webGetBytesWithStatus returns the body returned by an HTTP GET, as a []byte.
// It insists on a 200 response but sets the response code into status.
func webGetBytesWithStatus(url string, body *[]byte, status *int) error {
	return web.Get(url, web.ReadAllBody(body), web.Status(status))
}

// webGetBody returns the body returned by an HTTP GET, as a io.ReadCloser.
// It insists on a 200 response.
func webGetBody(url string, body *io.ReadCloser) error {
	return web.Get(url, web.Body(body))
}

// webGetBodyWithStatus returns the body returned by an HTTP GET, as a io.ReadCloser.
// It insists on a 200 response but sets the response code into status.
func webGetBodyWithStatus(url string, body *io.ReadCloser, status *int) error {
	return web.Get(url, web.Body(body), web.Status(status))
}
