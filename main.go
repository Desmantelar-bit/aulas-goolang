package main

import "fmt"

func main(){
	a, b := 10, 3
		fmt.Println("A soma é:", a+b)
		fmt.Println("A subtração é:", a-b)
		fmt.Println("A multiplicação é:", a*b)
		fmt.Println("A divisão é:", a/b, "e o resto é:", a%b)
	a++ // Pode-se fazer a++, a = a + 1 ou a += 1
		fmt.Println("Incrementar valor à a:", a)
	if a > 0 && b > 0 {
		fmt.Println("A e B são positivos")
	}
}