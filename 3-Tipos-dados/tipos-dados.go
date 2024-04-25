package main

import (
	"errors"
	"fmt"
)


func main() {

	var numero int64 = 10000000000000
	fmt.Println(numero)

	var numero2 int32 = 10000
	fmt.Println(numero2)
	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)
	//INT* = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	//Numeros reais
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 12333333333333333345.67

	fmt.Println(numeroReal1)
	fmt.Println(numeroReal2)

	//strings

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//char não existe, ele retorna um valor da tabela ask
	char := 'b'
	fmt.Println(char)

	/// fim de strinsgs

	texto := 5
	fmt.Println(texto)

	//boleano
	var booleano1 bool = true
	fmt.Print(booleano1)

	//error
	var erro error = errors.New("Erro interno")
	fmt.Print(erro)

}
