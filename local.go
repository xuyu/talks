// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var basePath string

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	httpListen := flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	flag.StringVar(&basePath, "base", pwd, "base path for slide template and static resources")
	flag.Parse()

	http.Handle("/static/", http.FileServer(http.Dir(basePath)))

	log.Printf("Open your web browser and visit http://%s/", *httpListen)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}
