package main

import ("fmt")

func main() {
	/*greating := "Hello, Bia!"
	fmt.Println(strings.Contains(greating, "Bia"))
	fmt.Println(strings.ReplaceAll(greating, "Hello", "Hi")) 
	fmt.Println(strings.ToUpper(greating))
	fmt.Println(strings.Index(greating, "Bia"))
	fmt.Println(strings.Split(greating, "Hello"))
	sort.Ints(ages)
	fmt.Println(ages)	
	index := sort.SearchInts(ages, 50)
	fmt.Println(index)
	sort.Strings(names)
	fmt.Println(names)
	index = sort.SearchStrings(names, "Sabrina")
	fmt.Println(index)*/
	
	/*ages := []int{50,80, 20, 100, 40}
	names := []string{"Bia","Sabrina", "Larissa"}
	x := 0
	for x< 5 {
		fmt.Println(x)
		x++
	}
	for i:= 0; i < 5; i++ {
		fmt.Println("for 2:", i)
	}

	for i:= 0; i <len(names); i++ {
		fmt.Println("for 3:", names[i])
	}

	for index, value := range names {
		fmt.Println("O índice é:", index,"e o valor é:", value)
	}

	for index, value := range ages {
		fmt.Println("O índice é:", index,"e o valor é:", value)
		ages = append(ages, 200)
		fmt.Println("O tamanho do vetor é:",len(ages), "A capacidade do vetor é:",cap(ages))

	}*/

	superherois := []string{"Superman", "Batman", "Flash"}

	for  index, value := range superherois {
			fmt.Println("O superheroi é:", value, "E sua posição é:", index)
			if value == "Batman" {
				fmt.Println("Gotham precisa de mim!")
			}
			if value == "Flash" {
				fmt.Println("Eu sou a velocidade!")
			}
			if value == "Superman" {
				fmt.Println("Eu sou um deus!")
			}
	}
}