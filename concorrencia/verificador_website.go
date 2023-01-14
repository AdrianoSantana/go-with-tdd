package concorrencia

type VerificadorWebsite func(string) bool
type Resultado struct {
	string
	bool
}

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan Resultado)
	for _, url := range urls {
		go func(u string) {
			canalResultado <- Resultado{u, vw(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}
	return resultados
}
