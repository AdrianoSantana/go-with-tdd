package soma

import (
	"reflect"
	"testing"
)

func TestSomaTodoResto(t *testing.T) {
	verificaSlices := func(t *testing.T, resultado []int, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado '%v' Esperado '%v'", resultado, esperado)
		}
	}

	t.Run("Soma slices", func(t *testing.T) {
		result := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSlices(t, result, esperado)
	})

	t.Run("Soma slices vázios", func(t *testing.T) {
		result := SomaTodoResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}

		verificaSlices(t, result, esperado)
	})

}
