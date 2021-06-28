// package render contains the rendering of templates to the ResponseWriter.
package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/andreasatle/bookings/config"
	"github.com/andreasatle/bookings/models"
)

// Container for parsed templates.
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tplCache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join("templates", "*.page.tmpl"))
	if err != nil {
		return tplCache, err
	}

	for _, page := range pages {
		// Of some reason, this has to be inside the for-loop
		baseTpl, err := template.New("baseTpl").Funcs(functions).ParseFiles(filepath.Join("templates", "base.layout.tmpl"))
		if err != nil {
			return tplCache, err
		}

		tplName := filepath.Base(page)
		log.Println("Parsing Template for Page:", page)
		tplCache[tplName], err = baseTpl.New(tplName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tplCache, err
		}
	}

	return tplCache, err
}

// RenderTemplate renders a template to the ResponseWriter.
func RenderTemplate(w http.ResponseWriter, page string, data *models.TemplateData) {

	var tplCache map[string]*template.Template
	if app.UseCache {
		tplCache = app.TemplateCache
		log.Println("Use template cache")
	} else {
		log.Println("Create template cache")
		var err error
		tplCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatalf("Error creating templates: %v\n", err)
		}
	}

	tpl, ok := tplCache[page]
	if !ok {
		log.Printf("no cached template <%s>", page)
		return
	}
	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("error writing to ResposeWriter: %v", err)
	}
}
