package main

import (
	"fmt"
	"strings"
)

// Função para repetir a string 'value' 'repetindo' vezes
func repetir(repetindo int, value string) string {
	str := strings.Repeat(value, repetindo)
	return str
}

func main() {
	fmt.Println(repetir(6, "oi"))
}


