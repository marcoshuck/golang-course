package main

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	db := connectDB() // This should return a connection handler, not a collection.
	defer db.Close()

	coll := db.Collection("blablabla")

	r := NewMongoRepository(coll)

	s := NewUserService(r)
	c := NewController(s)

	router.Get("/{uuid}", c.Get)

	http.ListenAndServe(":3030", router)
}
