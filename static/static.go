package static

import (
	"embed"
)

//go:embed resources/**
var EmbedStatic embed.FS
