package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo caso sejam nomeados
}

func main() {
	a, b := trocar(2, 3)
	fmt.Println(a, b)

	segundo, primeiro := trocar(7, 8)
	fmt.Println(segundo, primeiro)
}
