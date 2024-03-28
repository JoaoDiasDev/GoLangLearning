package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //Compilador conta

	for i, num := range numeros {
		fmt.Printf("%d) %d\n", i+1, num)
	}

	for _, num := range numeros {
		fmt.Printf("%d\n", num)
	}
}
