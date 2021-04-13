package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getCountHandler())

	http.ListenAndServe(":3030", nil)
}

func getCountHandler() http.HandlerFunc {
	var count int
	return func(w http.ResponseWriter, r *http.Request) {
		count += 1
		w.Write([]byte(fmt.Sprintf("Visitante actual: %d", count)))
		w.WriteHeader(http.StatusOK)
	}
}
