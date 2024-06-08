package auth

import (
	"errors"
	"net/http"
)

var users = map[string]string{
	"admin": "password",
}

func Authenticate(username, password string) error {
	if pass, ok := users[username]; ok && pass == password {
		return nil
	}
	return errors.New("invalid username or password")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if err := Authenticate(username, password); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
