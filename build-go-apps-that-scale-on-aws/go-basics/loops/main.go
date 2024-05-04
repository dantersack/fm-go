package main

import (
	"fmt"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "moose")
	// animals = slices.Delete(animals, 2, len(animals))

	// for i := 0; i < len(animals); i++ {
	// 	fmt.Printf("My animal is %s\n", animals[i])
	// }
	for idx, val := range animals {
		fmt.Printf("At position %d the animal is %s\n", idx, val)
	}

	// While loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
