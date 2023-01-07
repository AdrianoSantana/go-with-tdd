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
	testesArea := []struct {
		Forma    Forma
		Esperado float64
	}{
		{Forma: Retangulo{Altura: 4.0, Largura: 2.0}, Esperado: 8.0},
		{Forma: Circulo{Raio: 10}, Esperado: 314.1592653589793},
		{Forma: Triangulo{Altura: 5.0, Base: 2.0}, Esperado: 5.0},
	}

	for _, ta := range testesArea {
		resultado := ta.Forma.Area()
		if resultado != ta.Esperado {
			t.Errorf("Resultado: '%.2f' Esperado: '%.2f'", resultado, ta.Esperado)
		}
	}
}
