package dicionario

import "errors"

type Dictionary map[string]string

func (d Dictionary) Busca(dicionario map[string]string, palavra string) (string, error) {
	definicao, existe := dicionario[palavra]
	if !existe {
		return "", errors.New("Não foi possivel encontrar uma definição para essa palavra")
	}
	return definicao, nil
}
