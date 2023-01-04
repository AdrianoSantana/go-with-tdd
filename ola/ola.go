package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func ola(nome string) string {
	return prefixoOlaPortugues + nome
}
func main() {
	fmt.Println(ola("Adriano"))
}
