package main 

import "fmt"


//struct de usuarios 

type usuario struct {
	nome string
	idade uint8
	addres addres
}

type addres struct {
	logradouro string
	numero uint8
}

func main (){
 var u usuario
 u.nome = "Davi"
 u.idade = 21
 fmt.Println(u)

endereco := addres {"Rua estela pinto", 178} 

usuario2 := usuario {"Tiago", 31, endereco}
fmt.Println(usuario2)

usuario3 := usuario { nome: "Davi" , idade : 21  }
fmt.Println(usuario3)

}


/* package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco
	curso
}

type endereco struct {
	logradouro string
	numero     uint8
	bairro     string
	cidade     string
}

type curso struct {
	curso1 string
}

type

func main() {
	c := curso{"Eletricista"}
	e := endereco{"Rua estela pinto", 178, "Ponto chic", "Nova iguaçu"}
	u := usuario{"Tiago", 31, e, c}
	fmt.Println(u)
}/*



