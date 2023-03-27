package main

import "fmt"

func main() {

	var numero1 int
	var numero2 int

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scan(&numero1)

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		print("O número maior é ", numero1)
	} else if numero1 < numero2 {
		print("O maior número maior é:", numero2)
	} else {
		print("OS dois números são iguais")
	}
}

