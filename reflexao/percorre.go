package reflexao

import "reflect"

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}

func percorre(x interface{}, f func(string)) {
	valor := obtemValor(x)

	quantidadeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		f(valor.String())
	case reflect.Struct:
		quantidadeValores = valor.NumField()
		obtemCampo = valor.Field
	case reflect.Slice, reflect.Array:
		quantidadeValores = valor.Len()
		obtemCampo = valor.Index
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorre(valor.MapIndex(chave).Interface(), f)
		}
	}

	for i := 0; i < quantidadeValores; i++ {
		percorre(obtemCampo(i).Interface(), f)
	}
}
