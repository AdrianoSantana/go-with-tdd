package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if idioma == "espanhol" {
		return "Hola, " + nome
	}
	return prefixoOlaPortugues + nome
}
func main() {
	fmt.Println(ola("Adriano", ""))
}
