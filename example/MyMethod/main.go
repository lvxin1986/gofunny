package main

import (
"fmt"
"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) SetRadius1(r float64 )  {
	c.radius = r
}

func (c *Circle) SetRadius2(r float64 )  {
	(*c).radius = r
}


func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	c3 := Circle{10}
	c4 := Circle{25}
	c3.SetRadius1(12)
	c4.SetRadius2(12)
	fmt.Println("Area of c1 is: ", c3.radius)
	fmt.Println("Area of c2 is: ", c4.radius)
}
