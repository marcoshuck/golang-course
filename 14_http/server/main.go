package main

import (
	"net/http"
)

func main() {
	// r := mux.NewRouter()
	http.HandleFunc("/", route)
	http.ListenAndServe(":3030", nil)
}

// User -> Server -> Route -> Handler

// route is used by the standard HTTP library
func route(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.EscapedPath() == "/" {
		handler(w, r)
	}
}

// handler was used for the mux router example.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
	w.WriteHeader(http.StatusOK)
}
