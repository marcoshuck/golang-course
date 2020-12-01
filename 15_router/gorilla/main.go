package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var users []User

func init() {
	users = []User{
		{
			Username: "marcoshuck",
			Email:    "marcos@huck.com.ar",
			Password: "1234",
		},
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":8080", r)
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while marshalling slice of users"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
