package files

import (
	"embed"
)

//go:embed templates/*
var Templates embed.FS
