// +build dev

package main

import "net/http"

var Assets = http.Dir("assets")
