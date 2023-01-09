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
		comparaErro(t, err, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {
	t.Run("Insere uma palavra inexistente", func(t *testing.T) {
		dicionario := Dictionary{}
		err := dicionario.Adiciona("teste", "Isso é apenas um teste")
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, "teste", "Isso é apenas um teste")
	})

	t.Run("Insere uma palavra existente", func(t *testing.T) {
		dicionario := Dictionary{"teste": "Definição antiga"}
		err := dicionario.Adiciona("teste", "Nova definição")
		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, "teste", "Definição antiga")
	})
}

func TestAtualizacao(t *testing.T) {
	t.Run("Palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "Antiga definicao"
		dicionario := Dictionary{palavra: definicao}

		novaDefinicao := "Nova definicao"
		err := dicionario.Atualiza(palavra, novaDefinicao)
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("Palavra inexistente", func(t *testing.T) {
		palavra := "teste"
		definicao := "Teste definicao"
		dicionario := Dictionary{}
		err := dicionario.Atualiza(palavra, definicao)
		comparaErro(t, err, ErrPalavraInexistente)

	})
}

func TestDeletar(t *testing.T) {
	palavra := "any_palavra"
	definicao := "any_definicao"

	dicionario := Dictionary{palavra: definicao}
	dicionario.Deletar(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("Espera-se que '%s' seja deletado", palavra)
	}
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
