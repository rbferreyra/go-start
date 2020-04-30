package main

import (
	"fmt"
)

func main() {

	//temos o modelo tradicional para criar laço de repetição
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//aqui podemos incrementar a variável dentro do for
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
