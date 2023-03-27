package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um numero: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		print("Este número é par.")

	} else {
		print("Este número é impar")
	}
}
