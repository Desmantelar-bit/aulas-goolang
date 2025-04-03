package main

import "fmt"

func main() {
	var ages = [4]int{15, 15, 16, 16}
	nomes := [4]string{"Yan, Bruninha, Eduardinha, Pedrinha"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] ="Clark Kent"
	fmt.Println(nomes)
	//Slice
	var score = []int{100,200,300,400}
	fmt.Println(score)
	score[1] = 500
	fmt.Println(score, len(score), cap(score))
	rangeOne := score[1:3]
	fmt.Println(rangeOne)
	rangeTwo := score[2:]
	fmt.Println(rangeTwo)
	rangeThree := score[:3]
	fmt.Println(rangeThree)

	var superherois = []string{"Luffy", "Naruto", "Ichigo", "Goku"}
	fmt.Println(superherois)
	superherois = append(superherois, "Ben 10", "Superman", "Batman", "Flash")
	fmt.Println(superherois)
	fmt.Println(len(superherois), cap(superherois))
	
}