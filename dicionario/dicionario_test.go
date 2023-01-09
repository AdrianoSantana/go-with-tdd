package dicionario

import "testing"

func TestDicionario(t *testing.T) {
	dicionario := map[string]string{"teste": "Isso é apenas um teste"}

	resultado := Busca(dicionario, "teste")
	esperado := "Isso é apenas um teste"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
