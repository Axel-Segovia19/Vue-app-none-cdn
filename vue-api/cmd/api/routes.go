package main

import (
	"net/http"
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
	mux.Use(middleware.Recoverer)      // if things crash we will recover the application wont come to an end with recoverer
	mux.Use(cors.Handler(cors.Options{ //this allows what can be accessed
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

// this is linked to app.Login in the handlers
	mux.Post("/users/login", app.Login)
	// this is linked to the Logout handler
	mux.Post("/users/logout", app.Logout)

	mux.Route("/admin", func(mux chi.Router) { // in order for anyone to use routes inside this block they much have and authenticated token!
		mux.Use(app.AuthTokenMiddleware) // this is using AuthTokenMiddleware func to protect routes 

		mux.Post("/users", app.AllUsers) //this is calling the AllUsers handler to get all users in this protected route
		mux.Post("/users/save", app.EditUser) // allows us to edit and add a user
		mux.Post("/users/get/{id}", app.GetUser) // allows us to edit user by id
		mux.Post("/users/delete", app.DeleteUser) //* you call it by the route in the front end like so " fetch(process.env.VUE_API_URL + "/admin/users/delete"
	})



	return mux
}
