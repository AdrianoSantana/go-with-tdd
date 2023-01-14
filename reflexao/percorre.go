package reflexao

import "reflect"

func percorre(x interface{}, f func(string)) {
	valor := reflect.ValueOf(x)
	campo := valor.Field(0)
	f(campo.String())
}
