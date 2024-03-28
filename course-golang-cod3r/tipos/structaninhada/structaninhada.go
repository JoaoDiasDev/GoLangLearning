package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtd: 2, preco: 0.99},
			{produtoID: 2, qtd: 1, preco: 0.99},
			{produtoID: 3, qtd: 1, preco: 2.99},
		},
	}
	fmt.Println(pedido)
	fmt.Printf("Total do pedido: R$%.2f\n", pedido.valorTotal())
}
