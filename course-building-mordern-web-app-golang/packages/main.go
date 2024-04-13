package main

import (
	"log"

	"github.com/joaodiasdev/packagelearninggo/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "SomeType"
	myVar.TypeNumber = 5
	log.Println(myVar.TypeName, myVar.TypeNumber)
}
