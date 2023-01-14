package reflexao

import "testing"

func TestPercorrer(t *testing.T) {
	t.Run("Deve chamar uma função se o tipo da propriedade da struct for string", func(t *testing.T) {
		esperado := "Bia"
		var resultado []string

		x := struct {
			Nome string
		}{esperado}

		percorre(x, func(entrada string) {
			resultado = append(resultado, entrada)
		})

		if len(resultado) != 1 {
			t.Errorf("Número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
		}
	})
}
