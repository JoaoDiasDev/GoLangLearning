package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 3234.56,
			"Guga":    2234.56,
		},
		"J": {
			"João":   100234.56,
			"Joaquim": 5234.56,
		},
		"M": {
			"Marcelo": 3234.56,
			"Mariana": 2234.56,
		},
	}

	delete(funcsPorLetra, "M")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
