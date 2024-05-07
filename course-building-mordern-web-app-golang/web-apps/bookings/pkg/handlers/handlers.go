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

// Reservation renders the make a reservation page and displays form
func (repo *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (repo *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (repo *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (repo *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (repo *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contact.page.tmpl", &models.TemplateData{})
}
