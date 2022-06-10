package main

import "fmt"

func main() {

	var name string

	fmt.Println("Informe o seu nome: ")
	fmt.Scanf("%s", &name)

	fmt.Println("Olá " + name + "!")

}
