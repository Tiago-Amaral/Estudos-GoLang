package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

pessoa1 := pessoa {"Tiago","Lima", 20, 171}
fmt.Println(pessoa1)

estudante1 := estudante {pessoa1, "Engenheiro", "USP"}
fmt.Println(estudante1.altura)

}
