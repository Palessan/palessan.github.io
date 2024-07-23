package assets

import (
	"embed"
)

// wasmserve compatible
//
//go:embed *
var Assets embed.FS
