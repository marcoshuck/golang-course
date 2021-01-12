package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	// db, err := gorm.Open()
	// if err != nil {
	// 	log.Fatalf("Error while connecting to the database, error: %v", err)
	//	return
	// }
	//r := NewSQLRepository(db)

	r := NewInMemoryRepository(nil)

	s := NewUserService(r)
	c := NewController(s)

	router.Get("/{uuid}", c.Get)

	http.ListenAndServe(":3030", router)
}
