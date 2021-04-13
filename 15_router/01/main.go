package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(Greet))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type InputData struct {
	Name string
}

func Greet(w http.ResponseWriter, r *http.Request) {
	var b []byte
	buffer := bytes.NewBuffer(b)

	n, err := buffer.ReadFrom(r.Body)
	if err != nil {
		panic(err)
	}

	if n != r.ContentLength {
		panic(fmt.Errorf("content length is invalid"))
	}

	var data InputData
	err = json.Unmarshal(buffer.Bytes(), &data)
	if err != nil {
		panic(err)
	}

	auth := r.Header.Get("Authorization")
	var token string
	_, err = fmt.Sscanf(auth, "Bearer %s", &token)
	if err != nil {
		panic(err)
	}

	q := r.URL.Query()

	value := q.Get("page")
	var page int
	if page, err = strconv.Atoi(value); err != nil {
		panic(err)
	}

	value = q.Get("page_size")
	var pageSize int
	if page, err = strconv.Atoi(value); err != nil {
		panic(err)
	}

	List(page, pageSize)

	// chi: /books/{isbn}
	// gorilla/mux: /books/:isbn
	isbn := chi.URLParam(r, "isbn")

	var isbn int
	_, err = fmt.Sscanf(r.URL.RawPath, "/books/%d", &isbn)

	Get(isbn)
}

func Get(isbn string) {

}

func List(page int, size int) {

}
