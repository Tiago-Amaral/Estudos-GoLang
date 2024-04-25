package main 

import "fmt"



func imc (altura, peso float64)float64 {

	calculo := peso / (altura*altura)
	return calculo
}


func main (){

	total := imc (1.71, 90.0)
	imcFormatado := fmt.Sprintf("%.2f", total)

	if total < 16 {
		fmt.Println("magreza grave, seu imc é: ",imcFormatado )
	}else if total >= 16 && total <= 16.9{
		fmt.Println(" magreza moderada, seu imc é: ",imcFormatado )
	}else if total >= 17 && total <= 18.5{
		fmt.Println(" magreza leve, seu imc é: ",imcFormatado )
	}else if total >= 18.6 && total <= 24.9 {
		fmt.Println(" peso ideal, seu imc é: ",imcFormatado )
	}else if total >= 25 && total <= 29.9{
		fmt.Println(" sobrepeso, seu imc é: ", imcFormatado)
	}else if total >= 30 && total <= 34.9{
		fmt.Println(" obesidade grau I, seu imc é: ",imcFormatado )
	}else{
		fmt.Println(" obesidade grau III ou mórbida, seu imc é: ",imcFormatado )
	}

	
}