package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, esperado string, resultado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Esperado: '%s' resultado: '%s'", esperado, resultado)
		}
	}
	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		resultado := ola("Adriano")
		esperado := "Olá, Adriano"

		verificaMensagemCorreta(t, esperado, resultado)
	})

	t.Run("Diz 'olá, Mundo' se uma string vazia é passada", func(t *testing.T) {
		resultado := ola("")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, esperado, resultado)
	})
}
