package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// Write a response to the request
	respondWithJSON(w, 200, struct{}{})
}
