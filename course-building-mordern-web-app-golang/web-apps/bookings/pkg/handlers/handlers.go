package handlers

import (
	"net/http"

	"github.com/joaodiasdev/gobookings/pkg/config"
	"github.com/joaodiasdev/gobookings/pkg/models"
	"github.com/joaodiasdev/gobookings/pkg/render"
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
	remoteIP := request.RemoteAddr
	repo.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := repo.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to template
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
