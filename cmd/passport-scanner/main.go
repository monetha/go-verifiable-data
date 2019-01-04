package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/monetha/reputation-go-sdk/cmd"
	"github.com/shurcooL/httpgzip"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
)

func main() {
	flag.Parse()
	if cmd.PrintVersion() {
		return
	}

	log.Printf("listening on %q...", *listen)

	fileServer := httpgzip.FileServer(
		Assets,
		httpgzip.FileServerOptions{
			IndexHTML: true,
		},
	)

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
