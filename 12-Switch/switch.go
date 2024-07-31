package main

import "fmt"

func diaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Opção invalida"

	}

}

func diaSemana2(numero int)string{
switch {
case numero ==1:
	return "Domingo"
case numero ==2:
	return "Sehunda"
case numero ==3:
	return "Terça"
case numero ==4:
	return "Quarta"
case numero ==5:
	return "Quinta"
case numero ==6:
	return "Sexta"
case numero ==7:
	return "Sabado"
default:
	return "Invalido"


}
}

func main() {
	d := diaSemana(6)
	c := diaSemana2(3)
	fmt.Println(d,c)
}
