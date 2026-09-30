package main

import (
	"fmt"

	"github.com/ayush931/golang-learning/auth"
	"github.com/ayush931/golang-learning/user"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello world")
	auth.LoginWithCredentials("ayush", "Ayush@123")

	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "ayush@gmail.com",
		Name: "ayush",
	}
	fmt.Println(user)

	color.Red(user.Email)
	color.Yellow(user.Name)
}
