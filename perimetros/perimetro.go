package perimetros

type Retangulo struct {
	Altura  float64
	Largura float64
}

func Perimetro(rect Retangulo) (perimetro float64) {
	return 2 * (rect.Altura + rect.Largura)
}

func Area(rect Retangulo) float64 {
	return rect.Altura * rect.Largura
}
