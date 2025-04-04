package main

import ("fmt"; "sort")

func main() {
	/*greating := "Hello, Bia!"
	fmt.Println(strings.Contains(greating, "Bia"))
	fmt.Println(strings.ReplaceAll(greating, "Hello", "Hi")) 
	fmt.Println(strings.ToUpper(greating))
	fmt.Println(strings.Index(greating, "Bia"))
	fmt.Println(strings.Split(greating, "Hello"))
	*/
	ages := []int{50,80, 20, 100, 40}
	sort.Ints(ages)
	fmt.Println(ages)	
	index := sort.SearchInts(ages, 50)
	fmt.Println(index)
	names := []string{"Bia","Sabrina", "Larissa"}
	sort.Strings(names)
	fmt.Println(names)
	index = sort.SearchStrings(names, "Sabrina")
	fmt.Println(index)
}