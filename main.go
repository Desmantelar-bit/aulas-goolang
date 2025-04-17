package main

import ("fmt")

func dividir(dividendo float64, divisor float64) (float64, string){
	if divisor == 0 || dividendo == 0 {
		return 0, "Error: Division by zero"
	} else{
	result := dividendo / divisor
	return result, ""
	}	
}

// A função dividir recebe dois números e retorna o resultado da divisão e uma string de erro, se houver.
// A função retorna um número float64 e uma string.

func operacaoBasica(a int, b int) (int, int, int){
	soma:= a + b
	subtracao:= a - b
	multiplicacao:= a * b
	return soma, subtracao, multiplicacao
}
// A função operacaoBasica recebe dois números inteiros e retorna a soma, subtração e multiplicação deles.
// A função retorna três valores inteiros.

func main() { 
	resultado, erro := dividir(10, 2)
	if erro != "" {
		fmt.Println(erro)
	} else {
		fmt.Println("Resultado:", resultado,erro)
	}

	soma, subtracao, multiplicacao := operacaoBasica(10, 2)
	fmt.Println("Soma:", soma, "Subtração:", subtracao, "Multiplicação:", multiplicacao)
}