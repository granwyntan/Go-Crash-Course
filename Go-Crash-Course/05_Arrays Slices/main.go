package main

import "fmt"

func main() {
	// // Arrays - Of Fixed Length

	// // Declare
	// var fruitArray [2]string

	// // Assign Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// // Declare and Assign
	// fruitArray := [2]string{"Apple", "Orange"}

	// // Print
	// fmt.Println(fruitArray)
	// // Print with index
	// fmt.Println(fruitArray[1])

	// Slices - Length can vary/Change
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry", "Banana"}

	fmt.Println(fruitSlice)
	// Count
	fmt.Println(len(fruitSlice))

	// Certain items
	fmt.Println(fruitSlice[1:2]) // from index 1 (second) to right before index 2 (third)
	fmt.Println(fruitSlice[0:4]) // from index 0 (first) to right before index 4 (fifth)
}
