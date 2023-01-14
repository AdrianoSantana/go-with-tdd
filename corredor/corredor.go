package corredor

import (
	"fmt"
	"net/http"
	"time"
)

var limiteDezSegundos = 10 * time.Second

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func Corredor(a, b string) (vencedor string, err error) {
	return Configuravel(a, b, limiteDezSegundos)
}

func Configuravel(urlOne, urlTwo string, tempoLimite time.Duration) (vencedor string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("tempo limite excedido para %s e %s", urlOne, urlTwo)
	}
}
