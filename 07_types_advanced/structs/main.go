package main

import "fmt"

var (
	StatusOnline     Status = "online"
	StatusOffline    Status = "offline"
	StatusOnVacation Status = "on-vacation"
)

type Status string

// pointer receiver
func (s *Status) IsON() bool {
	return *s == StatusOnline
}

// value receiver      function
func (s Status) IsOFF() bool {
	return s == StatusOffline
}

func (s *Status) SetStatus(status Status) {
	*s = status
}

type User struct {
	Username string
	Password string
	Status   Status
}

func (u User) GetLinuxUsernamePassword() string {
	return fmt.Sprintf("%s@%s", u.Username, u.Password)
}

func main() {
	u := User{
		Username: "marcoshuck",
		Password: "1234",
		Status:   StatusOnline,
	}

	fmt.Println("Username:", u.Username)
	fmt.Println("Password:", u.Password)

	fmt.Println(u.Status.IsON())
	u.Status.SetStatus(StatusOnline)
	fmt.Println(u.Status.IsON())

	fmt.Println(u.GetLinuxUsernamePassword())
}
