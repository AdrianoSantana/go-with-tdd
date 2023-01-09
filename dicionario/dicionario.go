package dicionario

type Dictionary map[string]string

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

const (
	ErrNaoEncontrado    = ErrDicionario("Não foi possivel encontrar uma definição para essa palavra")
	ErrPalavraExistente = ErrDicionario("Já existe uma definição cadastrada para essa palavra")
)

func (d Dictionary) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dictionary) Adiciona(palavra string, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}
	return nil
}
