package main

import (
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	// Write a response to the request
	respondWithError(w, 400, "Something went wrong")
}
