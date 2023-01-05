package interacao

import "testing"

func TestIteracao(t *testing.T) {
	t.Run("Repetir 5 vezes", func(t *testing.T) {
		repeticoes := Repetir("a", 5)
		esperado := "aaaaa"

		if repeticoes != esperado {
			t.Errorf("Esperado '%s' Resultado '%s'", esperado, repeticoes)
		}
	})

	t.Run("Repetir x vezes", func(t *testing.T) {
		repeticoes := Repetir("a", 10)
		esperado := "aaaaaaaaaa"
		if repeticoes != esperado {
			t.Errorf("Esperado '%s' Resultado '%s'", esperado, repeticoes)
		}
	})
}
