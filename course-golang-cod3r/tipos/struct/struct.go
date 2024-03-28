package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	p1 := produto{
		nome:     "Notebook",
		preco:    1899.90,
		desconto: 0.10,
	}
	fmt.Println(p1, p1.precoComDesconto())
}
