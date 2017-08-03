package main

import "fmt"

func describe(thing interface{}) {
	fmt.Printf("%#v is of type %T\n", thing, thing)
}

func main() {
	var age int = 27
	percentage := 0.25 // float

	var name, neat = "Christian Roman", true

	describe(age)
	describe(percentage)
	describe(name)
	describe(neat)
}
