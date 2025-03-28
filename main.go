package main

import "fmt"

func main() {
	var senha int
	var usuario string
		fmt.Println("Digite a o seu login / usuário:")
		fmt.Scan(&usuario)
			if usuario == "admin" {
				fmt.Println("Digite a sua senha:")
				fmt.Scan(&senha)
					if senha == 1234 {
						fmt.Println("Acesso permitido!")
					} else {
						fmt.Println("Senha incorreta!")
					}
			} else {
				fmt.Println("Usuário incorreto!")
			}
}