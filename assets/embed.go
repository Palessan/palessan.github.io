package assets

import (
	"embed"
)

// file system cannot work with wasmserve

//go:embed *
var Assets embed.FS
