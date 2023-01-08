package carteira

import (
	"errors"
	"fmt"
)

var erroSaldoInsuficiente = errors.New("Não foi possivel concluir a transação: Saldo insuficiente")

type Bitcon float64

func (b Bitcon) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Carteira struct {
	saldo Bitcon
}

type Stringer interface {
	String() string
}

func (c *Carteira) Saldo() Bitcon {
	return c.saldo
}

func (c *Carteira) Depositar(valor Bitcon) {
	c.saldo += valor
}

func (c *Carteira) Retirar(valor Bitcon) error {
	if valor > c.saldo {
		return erroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}
