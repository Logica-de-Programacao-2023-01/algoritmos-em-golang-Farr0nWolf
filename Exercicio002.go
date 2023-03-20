package main

import "fmt"

func main() {

	var numero1 int
	var numero2 int
	var numero3 int

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scan(&numero1)

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scan(&numero2)

	fmt.Println("Digite o valor do terceiro número: ")
	fmt.Scan(&numero3)

	if numero1 < numero2 && numero1 < numero3 {
		print("O menor numero é:", numero1)

	} else if numero2 < numero1 && numero2 < numero3 {
		print("O menor número é:", numero2)

	} else if numero3 < numero1 && numero3 < numero2 {
		print("O menor número é:", numero3)

	} else if numero1 == numero2 && numero1 < numero3 {
		print("O menor número é:", numero1)

	} else if numero1 == numero3 && numero1 < numero2 {
		print("O menor numero é:", numero1)

	} else {
		print("Todos os números são iguais.")
	}

}
