package main 

import "fmt"

func main (){
//Ponteiro é uma referencia de memoria
// um ponteiro guarda um endereço de memoria de um valor inteiro

var variavel1 int = 10
var variavel2 int = variavel1
fmt.Println(variavel1, variavel2)

//Ponteiro
var variavel3 int
var ponteiro *int //ele mostra o endereço aonde esta guardado a variavel

variavel3 = 100
ponteiro = &variavel3 // precisa colocar o & para ele ser um ponteiro

fmt.Println(variavel3, ponteiro, *ponteiro) // quando eu coloco o * eu estou desreferenciando



}

