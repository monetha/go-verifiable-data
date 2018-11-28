// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"
	data "gitlab.com/monetha/reputation-go-sdk/cmd/passport-scanner"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}