package corredor

import (
	"net/http"
	"time"
)

func MedirTempoResposta(url string) time.Duration {
	inicio := time.Now()
	http.Get(url)
	return time.Since(inicio)
}

func Corredor(urlOne, urlTwo string) (vencedor string) {
	duracaoOne := MedirTempoResposta(urlOne)
	duracaoTwo := MedirTempoResposta(urlTwo)

	if duracaoOne < duracaoTwo {
		return urlOne
	}
	return urlTwo
}
