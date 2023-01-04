package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const holandes = "holandes"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaHolandes = "Hallo, "

func ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case holandes:
		prefixo = prefixoOlaHolandes
	}

	return prefixo + nome
}
func main() {
	fmt.Println(ola("Adriano", ""))
}
