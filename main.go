package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/axelboberg/lnkshrtnr/internal/redis"
)

type LinkRequest struct {
	URL string
}

type LinkResponse struct {
	ID string					`json:"id"`
}

// Get the parts of a path,
// useful when looking for 
// path parameters
func pathParts(p string) []string {
	return strings.Split(p, "/")
}

// A convenience function to write
// error messages with a status code
// back to the client
func error(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprint(w, message)
}

// Handle POST-requests to create new entries
// Will create a random key and assign the url
// from the request's body as its value
func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		error(w, 404, "404 page not found")
		return
	}

	var request LinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
			error(w, 400, "400 bad request")
			return
	}

	w.Header().Set("Content-Type", "application/json")

	var response LinkResponse
			response.ID = redis.SetRandom(request.URL)

	json.NewEncoder(w).Encode(response)
}

// Handle GET-requests by looking
// up the id provided in the url
// against the redis-store and
// make a 301 redirect
func get(w http.ResponseWriter, r *http.Request) {
	p := pathParts(r.URL.Path)
	if len(p) != 2 {
		error(w, 404, "404 page not found")
		return
	}

	id := p[1]
	url, ok := redis.Get(id)
	if !ok {
		error(w, 404, "404 page not found")
		return
	}

	log.Println("Redirecting to: " + url)
	http.Redirect(w, r, url, 301)
}

func main () {
	log.Println("Listening on port 3000")
	
	http.HandleFunc("/api/links", post)
	http.HandleFunc("/", get)
	http.ListenAndServe(":3000", nil)
}