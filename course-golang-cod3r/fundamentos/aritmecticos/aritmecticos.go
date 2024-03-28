package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", float32(a)/float32(b))
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // bitwise AND
	fmt.Println("Ou =>", a|b)  // bitwise OR
	fmt.Println("Xor =>", a^b) // bitwise XOR

	c := 3.0
	d := 2.0

	// outras operações...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Potência =>", math.Pow(c, d))

	// arredondamento
	fmt.Println("Arredondamento =>", int64(c))

	// raiz quadrada
	fmt.Println("Raiz quadrada =>", math.Sqrt(c))

}
