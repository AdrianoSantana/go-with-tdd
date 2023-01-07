package soma

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(numerosX ...[]int) []int {
	var somas []int
	for _, numeros := range numerosX {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

func SomaTodoResto(numerosX ...[]int) []int {
	var somas []int
	for _, numeros := range numerosX {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			resto := numeros[1:]
			somas = append(somas, Soma(resto))
		}
	}
	return somas
}
