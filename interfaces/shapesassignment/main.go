package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

type shape interface {
	//to meet the requirement for the shape interface
	//you need a function called getArea() that returns
	//type float64
	getArea() float64
}

//functions for the individual structs exist, but you also
//are creating a function for the shape interface
func getAea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	myTriangle := triangle{base: 5, height: 4}
	mySquare := square{sideLength: 4}

	fmt.Println("Triangle Area (should be 10)")
	// fmt.Println(myTriangle.getArea())
	myTriangle.getArea()
	fmt.Println("Square Area (should be 16)")
	// fmt.Println(mySquare.getArea())
	mySquare.getArea()
}
