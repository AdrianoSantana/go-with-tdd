package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("Deve verificar qual a url mais rápida", func(t *testing.T) {
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)
		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		urlLenta := servidorLento.URL
		urlRapida := servidorRapido.URL

		esperado := urlRapida
		resultado, _ := Corredor(urlLenta, urlRapida)

		if resultado != esperado {
			t.Errorf("Resultado '%s' Esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Retorna um erro se o servidor não responder dentro de 10 segundos", func(t *testing.T) {
		servidorA := criarServidorComAtraso(25 * time.Millisecond)
		servidorB := criarServidorComAtraso(25 * time.Millisecond)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Configuravel(servidorA.URL, servidorB.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("Esperava um erro, mas não obtive um.")
		}
	})

}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
