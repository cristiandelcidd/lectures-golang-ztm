package main

import "fmt"

func main() {
	var myName = "Cristian"
	fmt.Println("My name is", myName)

	var name string = "Mark"
	fmt.Println("name =", name)

	userRole := "admin"
	fmt.Println("The user's role is:", userRole)

	part1, part2 := 1, 2
	fmt.Println("Part 1 is:", part1, "Part 2 is:", part2)

	var sum int

	sum = 2 + 3
	fmt.Println("2 + 3 =", sum)

	var (
		firstName = "Cristian"
		surname   = "Delcid"
	)
	fmt.Println("My name is", firstName, surname)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
