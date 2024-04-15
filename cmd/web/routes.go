package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/vladgershun/daily-cents/ui"
)

func (app *application) routes() http.Handler {
	// Initialize a new multiplexer
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /user/signup", app.userSignUp)
	mux.HandleFunc("POST /user/signup", app.userSignUpPost)
	mux.HandleFunc("GET /user/login", app.userLogin)
	mux.HandleFunc("POST /user/login", app.userLoginPost)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
