package main

import "fmt"

func main() {
	//condicionais no go não difere muito das outras linguagens
	//mantém-se um padrão

	a := 10

	if a == 10 {
		fmt.Println("a == 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("nothing")
	}

	//pode-se atribuir uma variável na condicional
	//atribui o valor a variável, sendo que o b seja true
	//imprimindo a variável x
	b := false

	if x := "lorem ipsum"; b {
		fmt.Println(x)
	} else {
		fmt.Println("Nops")
	}
}
