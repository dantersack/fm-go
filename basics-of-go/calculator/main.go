package main

import (
	"fmt"
	"os"
)

func main() {
	var op, n1, n2 int

	PrintIntro()

	fmt.Scanf("%d", &op)
	fmt.Println("aber", op)

	if op != 1 && op != 2 && op != 3 && op != 4 {
		fmt.Println("Please select a valid operation")
		os.Exit(0)
	}

	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &n1)

	fmt.Println("Enter second number:")
	fmt.Scanf("%d", &n2)

	fmt.Print("Result: ")

	switch op {
	case 1:
		fmt.Println(n1 + n2)
	case 2:
		fmt.Println(n1 - n2)
	case 3:
		fmt.Println(n1 * n2)
	case 4:
		fmt.Println(n1 / n2)
	}
}
