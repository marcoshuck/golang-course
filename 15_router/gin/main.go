package main

import (
	"github.com/gin-gonic/gin"
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
	e := gin.New()

	e.GET("/users/{id}", getUsers)

	e.Run(":8080")
}

func getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &users)
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
