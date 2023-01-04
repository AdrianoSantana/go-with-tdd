package main

import "fmt"

func ola(nome string) string {
	return "Olá, " + nome
}
func main() {
	fmt.Println(ola("Adriano"))
}
