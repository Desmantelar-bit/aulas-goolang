package main

import ("fmt")

func erro() {
	fmt.Println("Erro: Algo deu errado! Tente novamente.")
	fmt.Println("Se o erro persistir, entre em contato com o suporte.")
}
func agradecer() {
	fmt.Println("Obrigado por usar nosso banco!")
	fmt.Println("Até logo!")
}

func depositar(saldo, dep, resultado float64) float64 {
	fmt.Println("Digite o valor a ser depositado:")
	fmt.Scan(&dep)
	if dep > 0 {
		fmt.Println("Depósito realizado com sucesso")
		resultado = saldo + dep
		fmt.Println("Seu saldo é: ", resultado)
		agradecer()
		return resultado

	} else if dep <= 0 {
		fmt.Println("Valor inválido")
		erro()
		return saldo
	}
	return saldo
}

func sacar(saldo, saque, resultado float64) float64 {
	fmt.Println("Digite o valor a ser sacado:")
	fmt.Scan(&saque)
	if saque > 0 && saque <= saldo {
		fmt.Println("Saque realizado com sucesso")
		resultado = saldo - saque
		fmt.Println("Seu saldo é: ", resultado)
		agradecer()
		return resultado

	} else if saque <= 0 {
		fmt.Println("Valor inválido")
		erro()
		return saldo
	
	} else if saque > saldo {
		fmt.Println("Saldo insuficiente, seu saldo é: ", saldo)
		erro()
		return saldo
	} 
	return saldo
}


func consultar(saldo float64) float64 {
	fmt.Println("Seu saldo é: ", saldo)
	return saldo
}

func saldação() {
	fmt.Println("Bem-vindo ao banco!")
	fmt.Println("Selecione uma opção:")
	fmt.Println("1 - Depositar")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Consultar saldo")
	fmt.Println("4 - Sair")
}

func main() {
	var saldo float64 = 1000.00
	var dep, saque, resultado float64
	var opcao int
	saldação()
	fmt.Scan(&opcao)
	for opcao != 4 {
		if opcao == 1 {
			depositar(saldo, dep, resultado)
			break
		} else if opcao == 2 {
			sacar(saldo, saque, resultado)
			break
		} else if opcao == 3 {
			saldo = consultar(saldo)	
			break	
		} else {
			fmt.Println("Opção inválida")
			main()
		}
	}
	for opcao == 4 {
		agradecer()
		break
	}
}