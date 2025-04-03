package main

import "fmt"

func main() {
	var ages = [4]int{15, 15, 16, 16}
	nomes := [4]string{"Yan, Bruninha, Eduardinha, Pedrinha"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] ="Clark Kent"
	fmt.Println(nomes)

}