package cumprimenta

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Santana")

	resultado := buffer.String()
	esperado := "Olá, Santana"

	if resultado != esperado {
		t.Errorf("Resultado: '%s' Esperado: '%s'", resultado, esperado)
	}
}
