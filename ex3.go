package main

import "fmt"

func main() {

	var idade int

	fmt.Print("Informe a sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade < 18 {
		fmt.Println("Você é menor de idade!")
	} else if idade >= 18 && idade <= 80 {
		fmt.Println("Você é maior de idade!")
	} else {
		fmt.Println("Você é maior de idade há um bom tempo!")
	}
}
