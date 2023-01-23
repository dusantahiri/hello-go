package handlers

import (
	"net/http"

	"github.com/dusantahiri/hello-go/pkg/config"
	"github.com/dusantahiri/hello-go/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func Home(response http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(response, "home.page.tmpl")
}

// About is the handler for the about page
func About(response http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(response, "about.page.tmpl")
}
