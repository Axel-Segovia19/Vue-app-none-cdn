package main

import (
	"net/http"
	"vue-api/cmd/api/internal/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes generates our routes and attaches them to handlers, using the chi router
// note that we return the type http.Handler, and not the *chi.Mux; since chi.Mux satisfies
// the interface requirements for http.Handler, it makes sense to return the type
// that is part of the standard library

// this allows us to reroute our handlers to not need inline handlers that then we can use to clean up our code
func (app *application) routes() http.Handler {
	mux := chi.NewRouter() 
	mux.Use(middleware.Recoverer) // if things crash we will recover the application wont come to an end with recoverer
	mux.Use(cors.Handler(cors.Options{ //this allows what can be accessed 
		AllowedOrigins:  []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Get("/users/login", app.Login) // this is linked to app.Login in the handlers
	mux.Post("/users/login", app.Login)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request){
		var users data.User
		all, err := users.GetAll()
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		app.writeJSON(w, http.StatusOK, all)
	})

	return mux
}