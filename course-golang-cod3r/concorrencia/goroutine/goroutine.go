package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Forma comum
	//	fale("Maria", "Pq vc não fala comigo?", 3)
	//  fale("João", "So posso falar depois de vc!", 1)

	// Forma concorrente goroutine
	// go fale("Maria", "Mari ei?", 500)
	// go fale("João", "João upa!", 500)

	go fale("Maria", "Pq vc não fala comigo?", 10)
	fale("João", "So posso falar depois de vc!", 5)
}
