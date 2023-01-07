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
