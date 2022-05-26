package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

//TODO this is a helper page to read and write json from this page alone that you can then reference too to keep code dry
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
		maxBytes := 1048576 // one megabyte limiting the max fiile size size that you will accept as a post
		r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))


		dec := json.NewDecoder(r.Body)
		err := dec.Decode(data)
		if err != nil {
			return err
		}

		err = dec.Decode(&struct{}{}) //want to make sure the body has a sinlge json value and not muliple
		if err != io.EOF { //end of File 
			return errors.New("body must have only a single json value")
		}

	return nil
}


func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")// converting go code into json
	if err != nil {
		return err
	}

	if len(headers) > 0 { //to make sure there is only 1 json value 
		for key, value := range headers[0] {
			w.Header()[key] = value 
		}
	}

	w.Header().Set("Content-Type", "application/json") //setting the content type into json 
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}


func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) { //this will write an error back to the client 
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	app.writeJSON(w, statusCode, payload) //wriiting to json the status and payload data if error is found 
}