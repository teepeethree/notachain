package main

import (
	"bufio"
	"log"
	"net/http"
	"notachain/routes"
	"os"
	"strings"
)

func loadEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
}

func main() {
	loadEnv()

	mode := os.Getenv("DEVELOPMENT_MODE")
	if mode == "dev" {
		routes.SetDevMode(true)
	}

	routes.LoadSharedTemplates(staticFiles)
	routes.LoadPageTemplates(staticFiles)

	registerStaticRoutes()
	registerRoutes()

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
