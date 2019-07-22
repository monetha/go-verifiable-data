// +build ignore

package main

import (
	"log"

	data "github.com/monetha/go-verifiable-data/cmd/passport-scanner/assets-data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
		Filename:     "assets-data/assets_vfsdata.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
