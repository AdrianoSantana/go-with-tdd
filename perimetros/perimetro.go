package perimetros

import "math"

type Forma interface {
	Area() float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

func Perimetro(rect Retangulo) (perimetro float64) {
	return 2 * (rect.Altura + rect.Largura)
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}
