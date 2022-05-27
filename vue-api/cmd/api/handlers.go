package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"vue-api/cmd/api/internal/data"

	"github.com/go-chi/chi/v5"
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
	
	// look up the user by email
	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}
	//validate the user's password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err!= nil || !validPassword {
		app.errorJSON(w, errors.New("invalid username/password"))
		return 
	}

	// make sure user is active
	if user.Active == 0 {
		app.errorJSON(w, errors.New("user is not active"))
		return 
	}

	// we have a valid user, so generate a token 
	token, err := app.models.Token.GenerateToken(user.ID, 24 *time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// save it to the database
	err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}


	//send back a response
	payload = jsonResponse{
		Error: false,
		Message: "logged in",
		Data: envelope{"token": token, "user": user},
	}

	err = app.writeJSON(w, http.StatusOK, payload) // writing to json if error is not found from write json helper
	if err != nil {
		app.errorLog.Println(err)
	}
}


func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
	app.errorJSON(w, errors.New("invalid json"))
	return
	}

	payload := jsonResponse{
		Error: false,
		Message: "logged out",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	var users data.User
	all, err := users.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	payload := jsonResponse{
		Error: false,
		Message: "success",
		Data: envelope{"users": all},
	}

	app.writeJSON(w, http.StatusOK, payload)
}
//allows us to edit and add a user
func (app *application) EditUser(w http.ResponseWriter, r *http.Request) { 
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if user.ID == 0 {
		// add user
		if _, err := app.models.User.Insert(user); err != nil {
			app.errorJSON(w, err)
		return
		}
	} else {
		//editing user
		u, err := app.models.User.GetOne(user.ID)
		if err != nil {
			app.errorJSON(w, err)
		return
	}

	u.Email = user.Email
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Active = user.Active

	if err := u.Update(); err != nil {
		app.errorJSON(w, err)
		return
	}

	// if password != string, update password
	if user.Password != "" {
		err := u.ResetPassword(user.Password)
		if err != nil {
				app.errorJSON(w, err)
			return
		}
	}
}

	payload := jsonResponse{
		Error: false,
		Message: "Changes saved",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) { //allows us to edit user by id
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user, err := app.models.User.GetOne(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, user)
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.User.DeleteByID(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error: false,
		Message: "User Deleted",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}