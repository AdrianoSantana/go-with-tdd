package reflexao

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func TestPercorrer(t *testing.T) {
	t.Run("Deve chamar uma função se o tipo da propriedade da struct for string", func(t *testing.T) {
		casos := []struct {
			Nome              string
			Entrada           interface{}
			ChamadasEsperadas []string
		}{
			{"Struct com um campo string", struct{ Nome string }{"Ana"}, []string{"Ana"}},
			{
				"Struct com mais de um campo string",
				struct {
					Nome      string
					Sobrenome string
				}{"Ana", "Santana"},
				[]string{"Ana", "Santana"},
			},
			{
				"Struct sem campo do tipo string",
				struct {
					Nome  string
					Idade int
				}{"Ana", 26},
				[]string{"Ana"},
			},
			{
				"Struct aninhada",
				Pessoa{
					"Ana",
					Perfil{26, "Recife"},
				},
				[]string{"Ana", "Recife"},
			},
			{
				"Ponteiros para coisas",
				&Pessoa{
					"Ana",
					Perfil{26, "Recife"},
				},
				[]string{"Ana", "Recife"},
			},
			{
				"Slices",
				[]Perfil{
					{33, "Interior"},
					{26, "Recife"},
				},
				[]string{"Interior", "Recife"},
			},
			{
				"Arrays",
				[2]Perfil{
					{33, "Londres"},
					{34, "Reykjavík"},
				},
				[]string{"Londres", "Reykjavík"},
			},
			{
				"Maps",
				map[string]string{
					"Ana": "Ana",
					"Bia": "aiB",
				},
				[]string{"Ana", "aiB"},
			},
		}

		for _, caso := range casos {
			var resultado []string
			percorre(caso.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})
			if len(resultado) != len(caso.ChamadasEsperadas) {
				t.Errorf("Número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
			}

			if !reflect.DeepEqual(resultado, caso.ChamadasEsperadas) {
				t.Errorf("Resultado: '%v', Esperado: '%v'", resultado, caso.ChamadasEsperadas)
			}
		}

	})
}
