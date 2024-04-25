 package main

 import (
	"fmt"
	"time"

 )

 func main () {
i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)

	}
	fmt.Println(i)

	for j:= 0; j < 10; j++ {
		fmt.Println("incrementanddo j", j)
		time.Sleep(time.Second )
	}

nomes := [3] string {"t", "n", "l"}

for indice, nome := range nomes {
	fmt.Println(indice, nome)
}
 }