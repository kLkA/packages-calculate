package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed public/*
var content embed.FS

func Handler() http.Handler {
	subFS, _ := fs.Sub(content, "public")
	return http.FileServer(http.FS(subFS))
}
