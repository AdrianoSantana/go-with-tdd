package adicionador

import "testing"

func TestAdicionadotr(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Esperado '%d' resultado '%d'", esperado, soma)
	}
}
