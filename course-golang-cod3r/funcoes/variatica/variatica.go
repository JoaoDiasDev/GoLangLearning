package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	if total == 0 {
		return 0
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média 1: %.2f ", media(7.7, 7.8, 7.9, 8.0, 8.1, 8.2, 8.3, 8.4, 8.5))
	fmt.Printf("Média 2: %.2f ", media(7.7, 7.9, 8.0, 8.1, 8.2, 8.3, 8.4, 8.5))
	fmt.Printf("Média 3: %.2f ", media())
}
