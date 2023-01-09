package dicionario

type Dictionary map[string]string

func (d Dictionary) Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}
