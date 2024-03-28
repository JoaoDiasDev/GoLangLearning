package main

import "math"

// Iniciando com letra maiúscula é publico (visível dentro e fora do pacote)

// Iniciando com letra minúscula é privado (visível apenas dentro do pacote)

// Por exemplo...
// Ponto - gerará algo publico
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

// catetos calcula o valor da hipotenusa
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
