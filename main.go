package main

import "fmt"

func main() {
	var numero1 float32
	var numero2 float32
	fmt.Println("Escreva o primeiro número:")
	fmt.Scanln(&numero1)
	fmt.Println("Escreva o segundo número:")
	fmt.Scanln(&numero2)
	fmt.Printf("Soma: %.2f\n", numero1+numero2)
	fmt.Printf("Subtração: %.2f\n", numero1-numero2)
	fmt.Printf("Multiplicação: %.2f\n", numero1*numero2)
	fmt.Printf("Divisão: %.2f\n", numero1/numero2)
}