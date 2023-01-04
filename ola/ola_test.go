package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola("Adriano")
	esperado := "Olá, Adriano"

	if resultado != esperado {
		t.Errorf("Esperado: %s resultado: %s", esperado, resultado)
	}

}
