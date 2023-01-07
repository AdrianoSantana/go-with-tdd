package carteira

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcon(10.0))
	resultado := carteira.Saldo()
	esperado := Bitcon(10.0)

	if resultado != esperado {
		t.Errorf("Resultado: '%s', Esperado: '%s'", resultado, esperado)
	}
}
