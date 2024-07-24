package main

import "fmt"

const prefixoPortugues = "Olá, "

func Ola(nome string) string {
	return prefixoPortugues + nome
}
func main() {
	fmt.Println(Ola("mundo"))
}




