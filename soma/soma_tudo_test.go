package soma

import (
	"reflect"
	"testing"
)

func TestSomaTudo(t *testing.T) {
	t.Run("Utiliza dois slices", func(t *testing.T) {
		resultado := SomaTudo([]int{1,2}, []int{0,9})
		esperado := []int{3,9}

		if !reflect.DeepEqual(esperado, resultado) {
			t.Errorf("Esperado '%d' Resultado '%d'", esperado, resultado)
		}
	})
}