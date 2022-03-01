package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	one, two, three := 1, 2, 3

	if one > two && one > three {
		fmt.Println("Number one is greater than two and three.")
	} else if two > one && two > three {
		fmt.Println("Number two is greater than one and three.")
	} else if three > one && three > two {
		fmt.Println("Number three is greater than one and two.")
	} else {
		fmt.Println("One or more numbers are equal.")
	}

	if average(one, two, three) > 0 {
		fmt.Println("Acceptable grades.")
	} else {
		fmt.Println("Instructor did a bad job.")
	}
}
