/*Zyon Morter 
1/30/23*/
package main

import (
	"fmt"
	"math"
)

// question 1
type triangle struct {
	base, height float64
}


//question 2 and 3 

func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}
func (t triangle) parameter() float64{
	return (t.base)*3
	
}

//question 6
type circle struct {
	radius float64
}

// question 7
func (c circle) area() float64 {
	const pi = 3.14159
	return (math.Pow( c.radius,2)) * pi
	 
}

// question 8
func (c circle) perimeter() float64 {
	const pi = 3.14159 
	return pi * c.radius * 2
	
}

func square(side float64) (float64, float64) {

	area := math.Pow(side, 2)
	perimeter := 4 * side
	return area, perimeter

}
// Helper function to compare, used in testing.
func withTolerance(a, b float64) bool {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}

func main () {
	

	//quesion 4 and 5
	var t triangle
fmt.Println(t.area(),t.parameter())

var c circle
fmt.Println(c.area(),c.perimeter())

fmt.Println(square(2))


}
