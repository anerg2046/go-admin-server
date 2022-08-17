package asset

import (
	"embed"
)

//go:embed views
var ViewFS embed.FS

//go:embed public
var PublicFS embed.FS
