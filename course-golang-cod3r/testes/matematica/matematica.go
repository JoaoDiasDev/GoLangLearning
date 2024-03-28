package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular a media de um conjunto de valores
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	if total == 0 {
		return 0
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
