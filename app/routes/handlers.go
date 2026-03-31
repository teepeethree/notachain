package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var templates map[string]*template.Template

var InlinedHeader string
var InlinedFooter string

func SetTemplates(t map[string]*template.Template) {
	templates = t
}

func SetInlinedTemplates(header, footer string) {
	InlinedHeader = header
	InlinedFooter = footer
}

func LoadSharedTemplates(fs embed.FS) {
	header := string(readFile(fs, "routes/_header.html"))
	footer := string(readFile(fs, "routes/_footer.html"))

	if !devMode {
		mainCSS := string(readFile(fs, "static/style.css"))
		header = strings.Replace(header,
			`<link rel="stylesheet" href="/style.css" />`,
			`<style>`+mainCSS+`</style>`, 1)

		appJS := string(readFile(fs, "static/app.js"))
		footer = strings.Replace(footer,
			`<script src="/app.js"></script>`,
			`<script>`+appJS+`</script>`, 1)
	}

	SetInlinedTemplates(header, footer)
}

func LoadPageTemplates(fs embed.FS) {
	t := make(map[string]*template.Template)

	for _, name := range []string{"index", "whitepaper", "token-paper", "registry"} {
		html := readFile(fs, "routes/"+name+".html")
		if html == nil {
			log.Printf("Failed to load %s", name)
			continue
		}
		full := InlinedHeader + string(html) + InlinedFooter
		t[name] = template.Must(template.New(name).Parse(full))
	}
	SetTemplates(t)
}

func readFile(fs embed.FS, path string) []byte {
	if devMode {
		b, _ := os.ReadFile(path)
		return b
	}
	b, _ := fs.ReadFile(path)
	return b
}

// servePage serves a templated page (header + content + footer).
// In dev mode it reads fresh from disk on every request.
func servePage(w http.ResponseWriter, name string) {
	w.Header().Set("Content-Type", "text/html")

	if devMode {
		h, _ := os.ReadFile("routes/_header.html")
		p, _ := os.ReadFile("routes/" + name + ".html")
		f, _ := os.ReadFile("routes/_footer.html")
		w.Write(h)
		w.Write(p)
		w.Write(f)
		return
	}

	if templates[name] == nil {
		http.Error(w, "page not available", 500)
		return
	}
	err := templates[name].Execute(w, nil)
	if err != nil {
		log.Printf("%s template error: %v", name, err)
	}
}

// serveRawPage serves a standalone page with no header/footer.
// In dev mode it reads fresh from disk on every request.
func serveRawPage(w http.ResponseWriter, name string) {
	w.Header().Set("Content-Type", "text/html")

	if devMode {
		p, _ := os.ReadFile("routes/" + name + ".html")
		w.Write(p)
		return
	}

	if templates[name] == nil {
		http.Error(w, "page not available", 500)
		return
	}
	err := templates[name].Execute(w, nil)
	if err != nil {
		log.Printf("%s template error: %v", name, err)
	}
}

var devMode bool

func SetDevMode(isDev bool) {
	log.Printf("Development mode is: %t", isDev)
	devMode = isDev
}

func GetDevMode() bool {
	return devMode
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, "index")
}

func WhitepaperHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, "whitepaper")
}

func TokenPaperHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, "token-paper")
}

func RegistryHandler(w http.ResponseWriter, r *http.Request) {
	servePage(w, "registry")
}
