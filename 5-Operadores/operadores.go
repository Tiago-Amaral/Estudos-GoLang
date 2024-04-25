package main

import "fmt"

func main() {
	//Aritimeticos
	s := 1 + 2
	sub := 5 - 1
	div := 2 / 3
	mul := 2 * 3
	rest := 10 % 2
	//Voce não pode  somar ou fazer qualquer conta com variaveis de tiposde dados diferentes como no exemplo a baixo
	//var numero1 int16 = 10
	//var numero2 int32 = 25

	//soma := numero1 + numero2

	//Operadores de atibuição

	var texto string = "String"
	text := "text"

	//operadores relacionais
	// 1 > 2 um é  maior que dois  /  false
	// 1 >= 2 um é maior ou igual a dois/  false
	// 1 == 2 um é igual a dois/ true
	// 1 <= 2 um é menor ou igual a dois/ true
	// 1 < 2 um é menor que dois/ true
	// 1 != 2 um é diferente que dois/ true

	//operadores logicos
	//&& e end
	// || ou our
	// ! not negação

	//Operador unario

	numero := 10
	numero++ // soma mais um
	numero += 15
	numero--    //diminui  um
	numero -= 2 // diminuir mais de um

	numero *= 3  // numero = numero *3
	numero /= 10 // numero = numero / 10
	numero %= 3  // numero = numero % 3

	//operador ternario
	//não existe em go, utilise o if e else

	var t string
	if numero > 5 {
		t = "Ele é maior que 5"
	} else {
		t = "Ele é menor que 5"
	}

	fmt.Println(t)

	fmt.Println(s, sub, div, mul, rest, texto, text, numero)
}
