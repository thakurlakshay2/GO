package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
	circumf() float64
}


type Circle struct {
	radius float64
}
type Square struct {
	length float64
}

//Square methods
func (s Square) area() float64{
	return s.length* s.length
}

func (s Square) circumf() float64{
	return s.length*4
}

//Circle Methods
func (s Circle) area() float64{
	return s.radius* s.radius*math.Pi
}

func (s Circle) circumf() float64{
	return 2*s.radius*math.Pi
}

func printShapeInfo(s Shape){
	fmt.Printf("area of %T is : %0.2f \n",s,s.area())
	fmt.Printf("Circumference of %T is : %0.2f \n",s,s.circumf())
}

func createData(){
	shapes := []Shape{
		Circle{radius:12.22},
		Circle{radius:10.22},
		Square{length:4},
		Square{length:6},
	}

	for _, v :=range shapes{
		printShapeInfo(v);
		fmt.Println("---")
	}
}