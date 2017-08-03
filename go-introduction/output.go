package main

import "fmt"

func main() {
	fmt.Println("This prints a newline after it")
	fmt.Printf("This doesn't\n")
	fmt.Printf("%s\n", "Format strings!")

	first_name := "Christian"
	last_name := "Roman"

	fmt.Println(fmt.Sprintf("%s %s\n", first_name, last_name))
}