// +build dev

package data

import "net/http"

var Assets = http.Dir("assets")
