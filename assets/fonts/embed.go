package fonts

import (
	_ "embed"
)

var (
	// no spaces allowed
	//go:embed Comic_Sans_MS.ttf
	Comic_Sans_MS []byte

	//go:embed Comic_Sans_MS_Bold.ttf
	Comic_Sans_MS_Bold []byte

	//go:embed press_start_2p.ttf
	Press_start_2p []byte

	//go:embed mplus_1p.ttf
	Mplus_1p []byte
)
