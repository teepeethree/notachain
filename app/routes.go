package main

import (
	"net/http"
	"notachain/routes"
	"strings"
)

func registerRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			routes.IndexHandler(w, r)
			return
		}
	})
	http.HandleFunc("/whitepaper", routes.WhitepaperHandler)
	http.HandleFunc("/economics", routes.EconomicsHandler)
	http.HandleFunc("/registry", routes.RegistryHandler)
	http.HandleFunc("/roadmap", routes.RoadmapHandler)
}

// RefCapture checks for ?ref= param and sets a cookie
func RefCapture(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ref := r.URL.Query().Get("ref")
		if ref != "" && len(ref) == 42 && strings.HasPrefix(ref, "0x") {
			http.SetCookie(w, &http.Cookie{
				Name:     "ref",
				Value:    ref,
				Path:     "/",
				MaxAge:   60 * 60 * 24 * 30, // 30 days
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})
		}
		next(w, r)
	}
}
