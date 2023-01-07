package soma

import (
	"reflect"
	"testing"
)

func TestSomaTodoResto(t *testing.T) {
	t.Run("Soma slices", func(t *testing.T) {
		result := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		if !reflect.DeepEqual(result, esperado) {
			t.Errorf("Resultado '%v' Esperado '%v'", result, esperado)
		}
	})

	t.Run("Soma slices vázios", func(t *testing.T) {
		result := SomaTodoResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}

		if !reflect.DeepEqual(result, esperado) {
			t.Errorf("Resultado '%v' Esperado '%v'", result, esperado)
		}
	})

}
