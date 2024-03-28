package main

import "fmt"

func main() {
	//mapas devem ser iniciados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[123456788] = "Pedro"
	aprovados[123456787] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456788])
	delete(aprovados, 123456788)
	fmt.Println(aprovados[123456788], "not found")
}
