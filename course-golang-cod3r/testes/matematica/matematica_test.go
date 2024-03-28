package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel()

	// Testing for a positive case
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}

	// Testing for a case with all zero values
	valorEsperado = 0
	valor = Media(0, 0, 0, 0)
	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
