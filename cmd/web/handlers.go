package main

import (
	"net/http"

	"github.com/vladgershun/daily-cents/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	mockBanks := []models.Bank{
		{ID: "1", Name: "IQCU"},
		{ID: "2", Name: "Chase"},
	}

	data := app.newTemplateData(r)
	data.Banks = mockBanks

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) userSignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Sign Up!"))
}

func (app *application) userSignUpPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Sign Up POST"))
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Login"))
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Login POST"))
}
