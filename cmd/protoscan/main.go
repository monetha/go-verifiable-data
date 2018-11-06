package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)

	fileServer := http.FileServer(Assets)
	err := http.ListenAndServe(*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		fileServer.ServeHTTP(resp, req)
	}))
	if err != nil {
		log.Fatalln(err)
	}
}
