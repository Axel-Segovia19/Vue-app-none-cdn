package main

import (
	"net/http"
)

//
type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

// Login is the handler used to attempt to log a user into the api
//this handler will be hit when we call it from out front end
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct { // what we get from the client/ front end we will use this to store the data received
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials    // var of creds of type credentials  holding email and password
	var payload jsonResponse // payload is of type jsonResponse that holds and error bool and a message string

	err := app.readJSON(w, r, &creds) // reading json from the readjson helper in helpers.go
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload) // writing json to http server if error found
	}

	//  TODO authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	//send back a response
	payload.Error = false
	payload.Message = "Signed in"

	err = app.writeJSON(w, http.StatusOK, payload) // writing to json if error is not found from write json helper
	if err != nil {
		app.errorLog.Println(err)
	}
}
