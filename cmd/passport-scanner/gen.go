package main

//go:generate env GOARCH=wasm GOOS=js go build -o assets/main.wasm web/main.go
//go:generate go run -tags=dev assets_generate.go
