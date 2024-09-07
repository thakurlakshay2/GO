package main

import "fmt"

func switchFunc() {
	val := 1

	switch val {
	case 1:
		fmt.Println("it is 1");
	case 2:
		fmt.Println("it is 2")
	default:
		fmt.Println("dhanki value daal")
	}
}