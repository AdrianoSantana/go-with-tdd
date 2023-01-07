package perimetros

import (
	"math"
	"testing"
)

func TestPerimetro(t *testing.T) {
	verificaTeste := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
		}
	}

	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0
	verificaTeste(t, resultado, esperado)
}

func TestArea(t *testing.T) {
	verificaTeste := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
		}
	}

	t.Run("Calcula Área de um Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0
		verificaTeste(t, resultado, esperado)
	})

	t.Run("Calcula área de um Círculo", func(t *testing.T) {
		circulo := Circulo{5}
		resultado := circulo.Area()
		esperado := math.Pi * (circulo.Raio * circulo.Raio)
		verificaTeste(t, resultado, esperado)
	})
}
