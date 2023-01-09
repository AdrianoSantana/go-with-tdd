package dicionario

import "testing"

func TestDicionario(t *testing.T) {
	dicionario := Dictionary{"teste": "Isso é apenas um teste"}
	t.Run("Busca uma palavra existente", func(t *testing.T) {
		resultado, _ := dicionario.Busca(dicionario, "teste")
		esperado := "Isso é apenas um teste"
		comparaString(t, resultado, esperado)
	})

	t.Run("Busca por uma palavra que não existe", func(t *testing.T) {
		dicionario := Dictionary{}
		_, err := dicionario.Busca(dicionario, "desconhecida")
		if err == nil {
			t.Fatal("Era esperado um erro!")
		}
		esperado := "Não foi possivel encontrar uma definição para essa palavra"
		comparaString(t, err.Error(), esperado)
	})
}

func comparaString(t *testing.T, resultado string, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
