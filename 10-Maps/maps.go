package main

import "fmt"



func main (){
usuario := map[string] string{
"nome" : "Tiago",
"Sobrenome" : "Lima",

}

usuario2 := map[string]map[string]string {
	"nome" : {
		"primeiro" : "RTiagho",
		"ultimo" : "Amaral",
	},
}
	fmt.Println(usuario, usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}