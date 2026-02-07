package main

import "fmt"

func main() {

	rect := Rectangle{width: 4, height: 5}
	circle := Circle{radius: 3}
	tri := Triangle{base: 6, height: 4}

	fmt.Println("=== BEFORE SCALE ===")

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	fmt.Printf("Triangle Area: %.2f\n", tri.Area())
	fmt.Printf("Triangle Perimeter: %.2f\n", tri.Perimeter())

	// Scale rectangle
	rect.Scale(2)

	fmt.Println("\n=== AFTER SCALE ===")

	fmt.Printf("Scaled Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Scaled Rectangle Perimeter: %.2f\n", rect.Perimeter())
}
