package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.ListenAndServe(":3030", r)
}

func route(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.EscapedPath() == "/" {
		handler(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
	w.WriteHeader(http.StatusOK)
}
