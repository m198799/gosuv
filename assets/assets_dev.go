//go:build !vfs
// +build !vfs

//go:generate go run assets_generate.go

package assets

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../res")
