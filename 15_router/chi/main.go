package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

var users []User

var count int

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
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(myFancyMiddleware)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", getUsers)
		r.Post("/", createUser)
	})

	http.ListenAndServe(":8080", r)
}

func myFancyMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Println("Count:", count)
		handler.ServeHTTP(w, r)
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {

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
