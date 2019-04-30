// +build tools

package tools

import (
	_ "github.com/alvaroloes/enumer"               // tool
	_ "github.com/ethereum/go-ethereum/cmd/abigen" // tool
	_ "golang.org/x/lint/golint"                   // tool
	_ "golang.org/x/tools/cmd/goimports"           // tool
	_ "honnef.co/go/tools/cmd/staticcheck"         // tool
)
