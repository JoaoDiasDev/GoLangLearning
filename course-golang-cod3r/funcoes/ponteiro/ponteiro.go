package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// desreferenciando (pegando o valor)
	*n++
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)
	inc2(&n) // por referencia
	fmt.Println(n)
}
