package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/dashboard", dashboard)
	mux.HandleFunc("/patients", patients)
	mux.HandleFunc("/patients/add", addPatient)
	mux.HandleFunc("/patients/view/", viewPatient)
	mux.HandleFunc("/patients/edit/", editPatient)
	mux.HandleFunc("/patients/delete/", deletePatient)
	mux.HandleFunc("/appointments", appointments)

	log.Print("Server listening on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
