package contagem

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct{}

func (s *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(w io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, ultimaPalavra)
}
