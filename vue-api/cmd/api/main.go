package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int //where we will hold all the configuration for this project
}

type application struct { //this will help to share bits of information with parts of you application
	config  config // share you configuration with the application
	infoLog *log.Logger //print info message to the console
	errorLog *log.Logger // print error messages to the console
}

func main() {
	var cfg config
	cfg.port = 8081 // listening to port 8080 on the front end so this one is for the back end

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // printing the date and time to the logger
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // log.Lshortfile will tell you where err took place 

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
	}

	// calls the func serve()
	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

// serve starts the web server 
func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr : fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}