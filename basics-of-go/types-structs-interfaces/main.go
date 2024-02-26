package main

import (
	"fmt"

	"dantejs.com/go/types-structs-interfaces/structs"
	"dantejs.com/go/types-structs-interfaces/types"
)

func main() {
	// test types
	location1 := types.Location("(35.5678, 30.1209)")
	location2 := types.Location("(-29.8765, -25.2537)")
	distance := location1.DistanceTo(location2)
	fmt.Println(distance)

	// test structs
	user := structs.User("dante", "dante@mail.com")
	userId := user.GetId()
	fmt.Println(user)
	fmt.Println(userId)
}
