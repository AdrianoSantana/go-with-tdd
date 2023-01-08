package carteira

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcon(10.0))
		confirmaSaldo(t, carteira, Bitcon(10.0))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcon(20)}
		err := carteira.Retirar(Bitcon(10.0))
		confirmaSaldo(t, carteira, Bitcon(10.0))
		confirmaErroInexistente(t, err)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcon(30)
		c := Carteira{saldo: saldoInicial}
		err := c.Retirar(100)
		confirmaSaldo(t, c, saldoInicial)
		confirmaErro(t, err, erroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcon) {
	t.Helper()
	resultado := carteira.Saldo()
	if resultado != esperado {
		t.Errorf("Resultado: '%s', Esperado: '%s'", resultado, esperado)
	}
}

func confirmaErro(t *testing.T, err error, esperado error) {
	t.Helper()
	if err == nil {
		t.Fatal("Esperava um erro, mas nada ocorreu!")
	}

	if err != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", err.Error(), esperado.Error())
	}
}

func confirmaErroInexistente(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Não esperava um erro")
	}
}
