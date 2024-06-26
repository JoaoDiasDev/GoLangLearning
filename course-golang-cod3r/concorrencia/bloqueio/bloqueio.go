package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Canal foi lido")
}

func main() {

	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Programa encerrado")
	fmt.Println(<-c) // deadlock
	fmt.Println("Nunca vai aparecer :(")
}
