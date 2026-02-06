package main

import "fmt"

func main() {
	r := Rectangle{Width: 5, Height: 10}
	c := Circle{Radius: 7}
	t := Triangle{Base: 6, Height: 8}

	fmt.Printf("Rectangle Area: %.2f\n", r.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", r.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", c.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", c.Perimeter())

	fmt.Printf("Triangle Area: %.2f\n", t.Area())
	fmt.Printf("Triangle Perimeter: %.2f\n", t.Perimeter())

	r.Scale(2)
	fmt.Printf("\nScaled Rectangle Area: %.2f\n", r.Area())
}
