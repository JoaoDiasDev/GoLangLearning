package main

import (
	generatorHtml "coder_golang/concorrencia/generator"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	urls := []string{url1, url2, url3}

	c1 := generatorHtml.Titulo(urls[0], urls[1], urls[2])

	// structure for concurrent control
	select {
	case t := <-c1:
		return t
	case <-time.After(10 * time.Second):
		return "Todos perderam"
	default:
		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.amazon.com",
	)
	fmt.Println(campeao)
}
