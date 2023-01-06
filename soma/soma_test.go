package soma

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1,2,3,4}
		resultado := Soma(numeros)
		esperado := 10
		if resultado != esperado {
			t.Errorf("Esperado '%v' Resultado '%v'", esperado, resultado)
		}
	})
}