package main

import "fmt"

//atribui o tipo do parâmetro, bem como tipo da saída
func funcName(a int) int {
	return a * a
}

//é possivel retorna a variável, bem como o tipo do mesmo
func namedReturn(a string) (x string) {
	x = a
	return
}

//é possível retorna várias saídas
func moreReturns(a string, b string) (string, string) {
	return a, b
}

//atribuindo ... no tipo, significa receber "múltiplos" parâmetros
func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}

	return res

}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(funcName(10))

	fmt.Println(namedReturn("Foobar"))

	x, y := moreReturns("Jhon", "Doe")
	fmt.Println(x, y)

	fmt.Println(variadicFunc(1, 2, 5, 10))

	//função anônima
	z := 0
	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(add())
}
