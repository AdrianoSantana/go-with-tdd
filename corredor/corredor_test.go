package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	servidorLento := criarServidorComAtraso(20 * time.Millisecond)
	servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

	urlLenta := servidorLento.URL
	urlRapida := servidorRapido.URL

	esperado := urlRapida
	resultado := Corredor(urlLenta, urlRapida)

	if resultado != esperado {
		t.Errorf("Resultado '%s' Esperado '%s'", resultado, esperado)
	}

	servidorLento.Close()
	servidorRapido.Close()
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
