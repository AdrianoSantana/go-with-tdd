package carteira

import "fmt"

type Bitcon float64

func (b Bitcon) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Carteira struct {
	bitcon Bitcon
}

type Stringer interface {
	String() string
}

func (c *Carteira) Saldo() Bitcon {
	return c.bitcon
}

func (c *Carteira) Depositar(valor Bitcon) {
	c.bitcon += valor
}
