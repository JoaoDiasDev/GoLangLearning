package handlers

import (
	"net/http"

	"github.com/joaodiasdev/hellowordwebapp/pkg/config"
	"github.com/joaodiasdev/hellowordwebapp/pkg/models"
	"github.com/joaodiasdev/hellowordwebapp/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(conf *config.AppConfig) *Repository {
	return &Repository{
		App: conf,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the home page handler
func (repo *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	// send the data to template
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
