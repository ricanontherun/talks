package main

import (
    "fmt"
    "errors"
)

// Multipy x by by return the product
func Multiply(x int, y int) int {
    return x * y
}

// Divide x by y, returning the quotient and the remainder.
func Divide(x int, y int) (int, int, error) {
    // Check for division by zero
    if y == 0 {
        return 0, 0, errors.New("Division by zero error!")
    }

    quotient := x / y;
    remainder := x % y

    // Notice I return nil for the 3rd variable.
    return quotient, remainder, nil
}

func main() {
    fmt.Printf("The result of %d x %d is %d\n", 10, 4,  Multiply(10, 4))

    // ok
    quotient, remainder, err := Divide(10, 4)
    if err == nil {
        fmt.Printf("The result of %d / %d is %d, remainder of %d\n", 10, 4, quotient, remainder)
    } else {
        fmt.Println(err)
    }

    // bad - division by zero
    if quotient, remainder, err := Divide(10, 0); err == nil {
        fmt.Printf("The result of %d / %d is %d, remainder of %d\n", 10, 0, quotient, remainder)
    } else {
        fmt.Println(err)
    }
}
