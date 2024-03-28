package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9, 10):
		return "A"
	case n.entre(8, 9):
		return "B"
	case n.entre(5, 8):
		return "C"
	case n.entre(3, 5):
		return "D"
	default:
		return "E"
	}
}

func main() {
	var n nota = 10
	fmt.Println(notaParaConceito(n))
	n = 8
	fmt.Println(notaParaConceito(n))
	n = 6
	fmt.Println(notaParaConceito(n))
	n = 4
	fmt.Println(notaParaConceito(n))
}
