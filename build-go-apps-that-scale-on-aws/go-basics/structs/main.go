package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// (person *Person) <-- This is called a "receiver"
func (person *Person) changeName(newName string) {
	fmt.Println("Address of person allocated memory:", &person.name)
	person.name = newName
}

func main() {
	var p = Person{
		name: "dante",
		age:  31,
	}

	fmt.Println("Address of p allocated memory:", &p.name)

	fmt.Printf("Name: %s - age: %d\n", p.name, p.age)
	p.changeName("sack")
	fmt.Printf("Name: %s - age: %d\n", p.name, p.age)

	p2 := newPerson("etnad", 13)
	fmt.Printf("Name: %s - age: %d\n", p2.name, p2.age)

	mySlice := []int{1, 2, 3}
	for idx := range mySlice {
		mySlice[idx]++
	}
	fmt.Println(mySlice)

}
