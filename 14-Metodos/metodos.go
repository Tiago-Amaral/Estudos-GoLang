package main 


import "fmt"


type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar(){
fmt.Printf("Salvando %s", u.nome)
}
func main (){
	usuario1 := usuario{"usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

}