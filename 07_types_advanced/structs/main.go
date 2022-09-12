package main

import "fmt"

// Are you tired from your original data types?
// string, integers, floats, whatever

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

var (
	StatusOnline     Status = "online"
	StatusOffline    Status = "offline"
	StatusOnVacation Status = "on-vacation"
	StatusAFK        Status = "afk"
)

// Status is your custom data type.
type Status string

type User struct {
	Username string
	Password string
	Status   Status
}

func main() {
	u := User{
		Username: "marcoshuck",
		Password: "1234",
		Status:   StatusOnVacation,
	}

	fmt.Println("Username:", u.Username)
	fmt.Println("Password:", u.Password)
	fmt.Println("Status:", u.Status)
}
