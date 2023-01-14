package reflexao

import "reflect"

func percorre(x interface{}, f func(string)) {
	valor := reflect.ValueOf(x)

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)
		if campo.Kind() == reflect.String {
			f(campo.String())
		}
	}
}
