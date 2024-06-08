package main

import (
	"html/template"
	"log"
	"net/http"
	"patient-management/internal/auth"
	"patient-management/internal/models"
	"strconv"
	"strings"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	ts, err := template.ParseFiles("./cmd/web/templates/base.tmpl", "./cmd/web/templates/"+tmpl)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "dashboard.tmpl", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		auth.Login(w, r)
		return
	}
	renderTemplate(w, "login.tmpl", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.tmpl", nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "dashboard.tmpl", nil)
}

func patients(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "patients.tmpl", patient)
}

func appointments(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "appointments.tmpl", nil)
}

var patient = []models.Patient{
	{ID: 1, FirstName: "Berk", LastName: "Toy", DOB: "23-01-1997", Address: "Salisbury-UK", Phone: "1111111111"},
}

func addPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := len(patient) + 1
		p := models.Patient{
			ID:        id,
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			DOB:       r.FormValue("dob"),
			Address:   r.FormValue("address"),
			Phone:     r.FormValue("phone"),
		}
		patient = append(patient, p)
		http.Redirect(w, r, "/patients", http.StatusSeeOther)
	} else {
		renderTemplate(w, "add_patient.tmpl", nil)
	}
}

func extractIDFromURL(r *http.Request) (int, error) {
	segments := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(segments[len(segments)-1])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func viewPatient(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromURL(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	for _, p := range patient {
		if p.ID == id {
			renderTemplate(w, "patient_detail.tmpl", p)
			return
		}
	}

	http.NotFound(w, r)
}

func editPatient(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromURL(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		for i, p := range patient {
			if p.ID == id {
				patient[i].FirstName = r.FormValue("first_name")
				patient[i].LastName = r.FormValue("last_name")
				patient[i].DOB = r.FormValue("dob")
				patient[i].Address = r.FormValue("address")
				patient[i].Phone = r.FormValue("phone")
				http.Redirect(w, r, "/patients", http.StatusSeeOther)
				return
			}
		}
	} else {
		for _, p := range patient {
			if p.ID == id {
				renderTemplate(w, "edit_patient.tmpl", p)
				return
			}
		}
		http.NotFound(w, r)
	}
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromURL(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	for i, p := range patient {
		if p.ID == id {
			patient = append(patient[:i], patient[i+1:]...)
			http.Redirect(w, r, "/patients", http.StatusSeeOther)
			return
		}
	}
	http.NotFound(w, r)
}
