package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		PhoneNumber: "123456789",
		Age:       30,
		BirthDate: time.Now(),
	}

	log.Println(user.FirstName,user.LastName, "Birth date:", user.BirthDate)
}

