package main

import (
	"fmt"
)

func main() {

	name := "Maria"

	switch name {
	case "Pedro":
		fmt.Println("Olá Pedro!")
	case "Maria":
		fmt.Println("Olá Maria!")
	default:
		fmt.Println("Não foi identificado o nome!")
	}
}
