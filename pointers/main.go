package main

import "fmt"

//metodo para realizar uma soma
//dessa forma vamos APONTAR a variável com um valor inteiro
//retornando a soma atribuido no ponteiro
func sum(a *int) int {
	*a = *a + 100
	return *a
}

func main() {

	b := 10
	fmt.Println(sum(&b)) //realiza a operação, endereçando a variável na memória
	fmt.Println(b)       // a variável antes era 10, passa a recuperar o valor endereçado na memória

	x := 10
	//endereça o valor da variável na memória, incluindo o "&" no inicio da variável
	fmt.Println(&x) // 0xc0000a4010

	//atribui o valor da variável y, com o endereçamento de memória da variável x
	y := &x
	fmt.Println(y) // 0xc0000a4010

	//para retorna o "valor original", basta incluir o "*" no inicio da variável
	fmt.Println(*y) //10

	//como a variável esta apontada para a mesma memória do x
	//quando atribui o valor a variavel y, a variável x é afetada
	*y = 20
	fmt.Println(x) //20

}
