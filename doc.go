// Copyright 2014 Benjamin Borbe. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package provide http-server helpers

Simple Web-Server:

	package main

	import (
		"github.com/bborbe/server"
		"github.com/bborbe/server/handler/static"
	)

	func main() {
		srv := server.NewServerPort(8080, static.NewHandlerStaticContent("Hello World"))
		srv.Run()
	}

*/
package server
