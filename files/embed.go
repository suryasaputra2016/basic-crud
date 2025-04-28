package files

import (
	"embed"
	"io/fs"
)

//go:embed templates/*
var Templates embed.FS

//go:embed assets/*
var assets embed.FS

func Assets() fs.FS {
	fs, _ := fs.Sub(assets, "assets")

	return fs
}
