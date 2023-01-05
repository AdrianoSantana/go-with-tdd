package interacao

func Repetir(valor string, qtdRepeticoes int) string {
	var repeticoes string
	for i := 0; i < qtdRepeticoes; i++ {
		repeticoes += valor
	}
	return repeticoes
}
