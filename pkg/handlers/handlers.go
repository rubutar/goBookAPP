package handlers

import (
	"net/http"

	"github.com/rubutar/goBookAPP/pkg/config"
	"github.com/rubutar/goBookAPP/pkg/models"
	"github.com/rubutar/goBookAPP/pkg/render"
)

// Repository used by the handlers
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "generals.page.html", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "majors.page.html", &models.TemplateData{})
}
func (m *Repository) Reserve(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "reserve.page.html", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, "availability.page.html", &models.TemplateData{})
}
