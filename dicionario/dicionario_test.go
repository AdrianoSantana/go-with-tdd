package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dictionary{"teste": "Isso é apenas um teste"}
	t.Run("Busca uma palavra existente", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "Isso é apenas um teste"
		comparaString(t, resultado, esperado)
	})

	t.Run("Busca por uma palavra que não existe", func(t *testing.T) {
		dicionario := Dictionary{}
		_, err := dicionario.Busca("desconhecida")
		if err == nil {
			t.Fatal("Era esperado um erro!")
		}
		comparaErro(t, err, errNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {
	dicionario := Dictionary{}
	dicionario.Adiciona("teste", "Isso é apenas um teste")

	comparaDefinicao(t, dicionario, "teste", "Isso é apenas um teste")
}

func comparaString(t *testing.T, resultado string, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func comparaDefinicao(t *testing.T, dicionario Dictionary, palavra, definicao string) {
	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("Não foi possivel encontrar palavra adicionada:", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
}
