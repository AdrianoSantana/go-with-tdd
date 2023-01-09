package dicionario

import "testing"

func TestDicionario(t *testing.T) {
	dicionario := Dictionary{"teste": "Isso é apenas um teste"}

	resultado := dicionario.Busca(dicionario, "teste")
	esperado := "Isso é apenas um teste"
	comparaString(t, resultado, esperado)
}

func comparaString(t *testing.T, resultado string, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
