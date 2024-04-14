package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
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
