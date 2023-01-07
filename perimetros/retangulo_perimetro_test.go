package perimetros

import "testing"

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
	retangulo := Retangulo{12.0, 6.0}
	resultado := Area(retangulo)
	esperado := 72.0
	if resultado != esperado {
		t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, esperado)
	}
}
