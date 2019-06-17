package main

import "fmt"

// Add two numbers
var x = 2
var y = 4
var result int

// AddTwoNumbers x + y
func AddTwoNumbers(x, y int) (result int) {
	result = x + y
	return result
}

func main() {
	result := AddTwoNumbers(x, y)
	fmt.Printf("The result is %d\n", result)
}
