package main

import (
	"fmt"
	"time"
)

func main() {

	soma := 0

	for i := 0; soma < 10; soma++ {
		fmt.Println(soma)
		soma += i
		//time.Sleep(2 * time.Second)
	}

	fmt.Println(soma)

	frutas := []string{"laranja", "maça", "banana", "uva"}
	for _, fruta := range frutas {

		fmt.Println(fruta)
		time.Sleep(2 * time.Second)
	}

}
