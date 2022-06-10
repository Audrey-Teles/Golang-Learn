package main

import "fmt"

func main() {

	var nome [5]string
	var altura [5]int

	for i := 0; i < 5; i++ {
		fmt.Println("------- PESSOA ", i+1, " -------")

		fmt.Println("Informe o seu nome: ")
		fmt.Scanf("%s", &nome[i])

		fmt.Println("Informe a sua altura: (em cm)")
		fmt.Scanf("%d", &altura[i])

	}

	for i := 0; i < 5; i++ {
		fmt.Println(nome[i], " tem ", altura[i], "cm de altura.")
	}

}
