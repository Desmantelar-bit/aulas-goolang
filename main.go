package main

import ("fmt")

func main() {
	var age int
	fmt.Println("Digite a sua idade:")
	fmt.Scan(&age)
	if age <=18 {
		fmt.Println("Você é menor idade.")
	} else if age >=18 && age <= 60 {
		fmt.Println("Você é adulto.")
	} else if age > 60 {
		fmt.Println("Você é idoso.")
	}
}