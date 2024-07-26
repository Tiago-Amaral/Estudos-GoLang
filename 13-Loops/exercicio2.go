package main

import "fmt"


func negative (x int) int {
	if x <=    0 {
		fmt.Println("já é negativo", x)
	
	return x
	}
	converter := x * (-1)
	return converter

	
	
}



func main(){



	fmt.Println(negative(-5))
}	
