package main

import (
	"embed"
	"net/http"
	"notachain/routes"
	"os"
)

//go:embed routes/** static/**
var staticFiles embed.FS

func serveStaticFile(path, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := readFiles(path)
		if data == nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Write(data)
	}
}

func registerStaticRoutes() {
	http.HandleFunc("/style.css", serveStaticFile("static/style.css", "text/css"))
	http.HandleFunc("/app.js", serveStaticFile("static/app.js", "application/javascript"))
	http.HandleFunc("/favicon-16x16.png", serveStaticFile("static/favicon-16x16.png", "image/png"))
	http.HandleFunc("/favicon-32x32.png", serveStaticFile("static/favicon-32x32.png", "image/png"))
	http.HandleFunc("/apple-touch-icon.png", serveStaticFile("static/apple-touch-icon.png", "image/png"))
	http.HandleFunc("/favicon.ico", serveStaticFile("static/favicon.ico", "image/x-icon"))
}

func readFiles(path string) []byte {
	if routes.GetDevMode() {
		b, _ := os.ReadFile(path)
		return b
	}
	b, _ := staticFiles.ReadFile(path)
	return b
}
