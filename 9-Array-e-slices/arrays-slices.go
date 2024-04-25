package main

import "fmt"

func main() {
	// o array tem um tamanho definido
	var array1 [5]string
	array1[0] = "P1" //adicionando um intem no array

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"} // array

	// um slice não tem um tamanho definido, ele é elastico
	slice := []int{10, 40, 05, 30, 60} // slice

	slice = append(slice, 18) // aqui está adicionando um novo intem no slice

	slice2 := array2[1:3] // aqui está criando um slice apastir de um array que já existe

	array2[1] = "p2 alterado" // aqui quando o array foi alterado,  o slice tambem foi, eele funciona como um apontador

	slice4 := [] int {12,45,56,76,89}

	array4 := [5] int {12,45,34,56}

	fmt.Println(array1, array2, slice, slice2, slice4, array4)

}
