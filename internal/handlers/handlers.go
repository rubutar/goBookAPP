package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rubutar/goBookAPP/internal/config"
	"github.com/rubutar/goBookAPP/internal/forms"
	"github.com/rubutar/goBookAPP/internal/models"
	"github.com/rubutar/goBookAPP/internal/render"
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}
func (m *Repository) Reserve(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reserve.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handles the posting of reservation form
func (m *Repository) PostReserve(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println((err))
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)

	form.Has("first_name", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "reserve.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	//send the data to the template
	render.RenderTemplate(w, r, "availability.page.html", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("%s to %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handler to request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
