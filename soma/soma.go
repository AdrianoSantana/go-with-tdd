package soma

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}