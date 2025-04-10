package main

import ("fmt")

func main() {
	var saldo float64 
	var opcao int
	var depósito float64
	var saque float64
	var resultado float64
	fmt.Println("Bem-vindo ao banco!")
	fmt.Println("Digite o valor inicial do saldo:")
	fmt.Scan(&saldo)
	fmt.Println("O que você deseja fazer?")
	fmt.Println("1 - Depositar")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Ver saldo")
	fmt.Scan(&opcao)
	switch opcao {
	case 1:
		fmt.Println("Digite o valor do depósito:")
		fmt.Scan(&depósito)
			if depósito <= 0{
				fmt.Println("Valor inválido!")
			} else {
				resultado = saldo + depósito
				fmt.Print("Depósito de ", depósito," realizado com sucesso! Saldo atual: ",resultado)
			}
		break
	case 2:
		fmt.Println("Digite o valor do saque:")
		fmt.Scanln(&saque)
			if saque > saldo {
				fmt.Println("Saldo insuficiente!")
			} else if saque <= 0 {
				fmt.Println("Valor inválido!")
			}else{
				resultado = saldo - saque
				fmt.Print("Saque de ", saque," realizado com sucesso! Saldo atual: ",resultado)
			}
		break
	case 3:
		fmt.Print("Seu saldo é: ",saldo)
		break
	default:
		fmt.Println("Opção inválida!")
		return

}
	fmt.Println("Obrigado por usar nosso banco!")
	fmt.Println("Até logo! Qualquer problema entre em contato com o suporte!")
}