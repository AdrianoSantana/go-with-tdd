package contagem

import (
	"bytes"
	"reflect"
	"testing"
)

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas += 1
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

const pausa = "pausa"
const escrita = "escrita"

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func TestContagem(t *testing.T) {
	t.Run("Imprimite de 3 até 1 e vai", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Contagem(buffer, &SpyContagemOperacoes{})
		resultado := buffer.String()
		esperado := `3
2
1
Vai!`

		if resultado != esperado {
			t.Errorf("Resultado: '%s', Esperado: '%s'", resultado, esperado)
		}
	})

	t.Run("Pausa antes de cada escrita", func(t *testing.T) {
		spy := &SpyContagemOperacoes{}
		Contagem(spy, spy)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spy.Chamadas) {
			t.Errorf("Esperado: %v, Resultado: %v", esperado, spy.Chamadas)
		}

	})
}
