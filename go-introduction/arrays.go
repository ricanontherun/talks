package main

import "fmt"

func main() {
	// arrays are fixed length
	var ids [10]int

	// slices are not. This gets weird.
	var ids_slice []int = ids[0:5]

	fmt.Println(ids)	
	fmt.Println(ids_slice)

	// Slices are like references to their underlying array.
	// modifying the first element of this slice modifies
	// the first element of the ids array.
	ids_slice[0] = 100

	fmt.Println(ids)	
	fmt.Println(ids_slice)

	var nil_slice []int

	initialized_slice := []bool{true, true, false}
	initialized_slice = append(initialized_slice, false)

	fmt.Println(nil_slice)
	fmt.Println(initialized_slice)
}

