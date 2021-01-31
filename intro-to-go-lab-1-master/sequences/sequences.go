package main

import (
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i := range slice {
		slice[i] = f(slice[i])
	}
}

func mapArray(f func(a int) int, array [5]int) {
	for i := range array {
		array[i] = f(array[i])
	}
	// fmt.Println(array) //works here but not in main
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	// mapSlice(addOne, intsSlice)
	// fmt.Println("map slice",intsSlice)

	// intsArray := [5]int{1, 2, 3, 4, 5}
	// mapArray(square, intsArray)
	// fmt.Println("map Array",intsArray) //stays the sameT

	// newSlice := intsSlice[1:3]
	// mapSlice(square, newSlice)
	// fmt.Println("new slice", newSlice)

	intsSlice = double(intsSlice)
	fmt.Println("double slice", intsSlice)

}
