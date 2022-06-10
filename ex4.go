package main

import "fmt"

func main() {

	var altura [5]float
	var nome [5]string

	for i := 0; i < 5; i++ {
		fmt.Println("Informe a sua altura:")
		fmt.Scanf("%f", &altura)
	}

}
