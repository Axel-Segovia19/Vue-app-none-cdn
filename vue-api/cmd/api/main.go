package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vue-api/cmd/api/internal/data"
	"vue-api/cmd/api/internal/driver"
)

// config is the type for all the application configuration
type config struct {
	port int //what port we want the web server to listen on 
}

//application is the type for all data we want to share with the various parts of our application we will
// share this information in most cases by using this type as the receiver for functions
type application struct { //this will help to share bits of information with parts of your application
	config  config // share your configuration with the application
	infoLog *log.Logger //print info message to the console
	errorLog *log.Logger // print error messages to the console
	models data.Models
}

// main is the main entry point for our application
func main() {
	var cfg config // variable of type config from the struct that holds a port of type int
	cfg.port = 8081 // listening to port 8080 on the front end so this one is for the back end

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // printing the date and time to the logger
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // log.Lshortfile will tell you where err took place 

	dsn := os.Getenv("DSN")
	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close() //will take care of closing connections still running by closing when program stops

	app := &application{ // this is a variable called app reference to the application struct
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		models: data.New(db.SQL),
	}


	// calls the func serve()
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

// serve() starts the web server 
//app * application is the receiver the functions serve() and returns an error
func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port) //app.config.port = port 8081 which is in the receiver which pulls from the declared cfg

	srv := &http.Server{
		Addr : fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(), //this is pointing to routes in handlers.go allowing us to use external handlers
	}

	return srv.ListenAndServe() // we can do this because listen and serve is a method attachd to that type
}