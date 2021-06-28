package handlers

import (
	"log"
	"net/http"

	"github.com/andreasatle/bookings/config"
	"github.com/andreasatle/bookings/models"
	"github.com/andreasatle/bookings/render"
)

// Repo contains the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo returns a new Repository instance.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the module pointer Repo.
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the handler for /
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	log.Println("Calling Home handler from", remoteIP)

	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for /about
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling About handler")

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{
			"text":      "Some text",
			"remote_ip": remoteIP,
		},
	})
}
