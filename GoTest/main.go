package main

import "fmt"

type CalculateArea interface {
	area() int
	perimeter() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length  int
	breadth int
}

func CalculateAreaInterface(obj CalculateArea) {
	fmt.Println(obj.area())
	fmt.Println(obj.perimeter())

}

func (s Square) perimeter() int {
	return 4 * s.side
}

func (s Square) area() int {
	return 4 * s.side
}

func (r Rectangle) perimeter() int {
	return r.length * r.breadth
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	s := Square{
		side: 4,
	}
	r := Rectangle{
		length:  5,
		breadth: 6,
	}
	CalculateAreaInterface(s)
	CalculateAreaInterface(r)

}
