package reflexao

func percorre(x interface{}, f func(string)) {
	f("any string")
}
