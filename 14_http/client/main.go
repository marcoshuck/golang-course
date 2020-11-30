package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"login"`
	Avatar   string `json:"avatar_url"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Blog     string `json:"blog"`
	Bio      string `json:"bio"`
}

func main() {
	res, err := http.Get("https://api.github.com/users/marcoshuck")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	var b []byte

	n, err := res.Body.Read(b) // This doesn't work
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Read: [%d] Data: [%s]\n", n, b)

	scanner := bufio.NewScanner(res.Body) // This works
	defer res.Body.Close()

	for scanner.Scan() {
		scanner.Text()
		b = append(b, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Request: %s\n", string(b))
	fmt.Printf("User: %+v\n", user)
}
