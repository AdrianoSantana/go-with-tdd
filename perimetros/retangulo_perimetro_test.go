package perimetros

import (
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
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()
		if resultado != esperado {
			t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
		}
	}

	t.Run("Calcula Área de um Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		esperado := 72.0
		verificaArea(t, retangulo, esperado)
	})

	t.Run("Calcula área de um Círculo", func(t *testing.T) {
		circulo := Circulo{10}
		esperado := 314.1592653589793
		verificaArea(t, circulo, esperado)
	})
}
