package perimetros

import "testing"

func TestPerimetro(t *testing.T) {
	verificaTeste := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
		}
	}

	resultado := Perimetro(10.0, 10.0)
	esperado := 40.0
	verificaTeste(t, resultado, esperado)
}

func TestArea(t *testing.T) {
	resultado := Area(12.0, 6.0)
	esperado := 72.0
	if resultado != esperado {
		t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
	}
}
