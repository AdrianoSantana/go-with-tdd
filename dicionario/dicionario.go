package dicionario

import "errors"

type Dictionary map[string]string

var errNaoEncontrado = errors.New("Não foi possivel encontrar uma definição para essa palavra")

func (d Dictionary) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", errNaoEncontrado
	}
	return definicao, nil
}

func (d Dictionary) Adiciona(palavra string, definicao string) {
	d[palavra] = definicao
}
