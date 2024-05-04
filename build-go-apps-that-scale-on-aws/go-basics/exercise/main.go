package main

import (
	"fmt"
	"math"
)

const PI = 3.14

type Circle struct {
	Radius float64
}

func (c Circle) Circumference() float64 {
	return 2 * PI * c.Radius
}

func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func main() {
	var circle = Circle{Radius: 3}
	fmt.Println(circle.Circumference())
	fmt.Println(circle.Area())
}
