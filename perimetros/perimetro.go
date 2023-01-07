package perimetros

func Perimetro(largura, altura float64) (perimetro float64) {
	return 2 * (largura + altura)
}

func Area(largura, altura float64) float64 {
	return largura * altura
}
