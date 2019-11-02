package main

import (
	"fmt"
	"math"
)


type Cycle struct {
	radius float64
}
type Rectangle struct {
	width,height float64
}
func (r Rectangle) area() float64{
	return r.width * r.height
}

func (c Cycle)area() float64  {
	return c.radius * c.radius *math.Pi
}

func main() {
	r1 := Rectangle{12,2}
	r2 := Rectangle{9,4}
	c1 := Cycle{10}
	c2 :=Cycle{25}
	fmt.Println("Area of rectangle is:",r1.area())
	fmt.Println("Area of rectangle is:",r2.area())
	fmt.Println("Area of Cycle is:",c1.area())
	fmt.Println("Area of Cycle is:",c2.area())
}