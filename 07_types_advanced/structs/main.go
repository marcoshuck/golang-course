package main

import "fmt"

type User struct {
	Username string
	Password string
}

func main() {
	u := User{
		Username: "marcoshuck",
		Password: "1234",
	}

	fmt.Println("Username:", u.Username)
	fmt.Println("Password:", u.Password)
}
