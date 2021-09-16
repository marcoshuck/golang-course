package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Car struct {
	Kilometers uint `json:"kilometers"`
}

type Cup struct {
	Volume float64 `json:"volume"`
	Color  string  `json:"color"`
	Owner  User
}

func NewCarA() Car {
	return Car{
		Kilometers: 200,
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Delivered   time.Time `json:"delivered"`
	Owner       User      `json:"owner"`
}

func main() {
	// 0 - 1024 (80)
	// 1025 -> 65355
	// 8080
	// 8081
	// 3000

	const port = 3000
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, http.HandlerFunc(handler))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Shutting server down...")
}

func handler(w http.ResponseWriter, r *http.Request) {
	printLog(r)
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodDelete:
		remove(w, r)
	}
}

// HEADER
// GET HTTP/1.1
//
// BODY
// { }
// ""
func remove(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Something has been deleted"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func post(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Something has been created"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func get(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	// http://localhost:3000/?a=1&b=2
	// Here is something for you: 3
	a := values.Get("a")
	b := values.Get("b")

	valueA, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	valueB, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := valueA + valueB

	_, err = w.Write([]byte(fmt.Sprintf("Here is something for you: %d", result)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func printLog(r *http.Request) {
	log.Printf("| %s | %s | %s | URL: %s\n", r.Method, r.Proto, r.UserAgent(), r.URL.String())
}
