package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João": 1234.56,
		"Maria": 4567.89,
		"Pedro": 7890.12,
	}

	funcsESalarios["Rafael"] = 9876.54
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
