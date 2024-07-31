package main 

import "fmt"

func dias(dia int)string{
	switch dia{
	case 1 :
		return "Domingo"
	case 2:
		return "segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	default:
		return "Opção invalida"
	}


}

func main() {
	d := dias(4)
	fmt.Println(d)
}
