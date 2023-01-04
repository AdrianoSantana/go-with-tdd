package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const holandes = "holandes"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaHolandes = "Hallo, "

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case holandes:
		prefixo = prefixoOlaHolandes
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoDeSaudacao(idioma) + nome
}
func main() {
	fmt.Println(ola("Adriano", ""))
}
