package main

import (
	"fmt"
)

// Write a program that calculates the year
// using a provided date of birth and age.
// HINT: Get the date of birth and age from stdin!
func calcYearBorn() {
	var birth_year int
	var age int

	fmt.Println("Enter your year of birth: ")
	fmt.Scanln(&birth_year)

	fmt.Println("Enter your current age: ")
	fmt.Scanln(&age)

	year := birth_year + age

	fmt.Println("\nThe year is ", year)

}

// Write a program that calculates the average weight of 5 people.
func calcAvgWeight() {
	// declaring an array of values
	weightArr := []int{100, 200, 300, 400, 500}

	// size of the array ie number of people
	n := 5

	sum := 0

	// loop through the array to get weights
	for i := 0; i < n; i++ {
		// adding the weights in array to the variable sum
		sum += (weightArr[i])
	}

	avg := (float64(sum)) / (float64(n))

	fmt.Println("\nAverage Weight of 5 people is = ", avg)

}

func main() {
	calcYearBorn()
	calcAvgWeight()
}
