package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func somaMatematica(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subt := n1 - n2
	return soma, subt
}

func subtra (n5, n6 int8) int8 {
  
	result := n5 - n6 
	return result
}

//caso eu queira descartar uma variavel preciso substitui-la por um "_" no lugar ddela
func main() {
    subt := subtra (20, 10)
	fmt.Println(subt)


	resultadoCalculosoma, resultadoCalculossubt := somaMatematica(10, 15)
	fmt.Println(resultadoCalculosoma, resultadoCalculossubt)

}
