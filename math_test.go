package main

import (
	"testing"
)


//call the triangle area test 
func TestTriangleArea(t *testing.T) {
	
	t1 := triangle{
		base:   3.5,
		height: 6,
	}
	got := t1.area()

	//what i got and what i expect the answer to be 

	expected := 10.5

	//comparing the resuults 
	equal := withTolerance(got, expected)

	if equal != true {
		t.Errorf("got %f, expected %f", got, expected)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	
	t1 := triangle{
		base:   3.5,
		height: 6,
	}
	got := t1.parameter()

	expected := 10.5

	equal := withTolerance(got, expected)

	if equal != true {
		t.Errorf("got %f, expected %f", got, expected)
	}
}

func TestCircleArea(t *testing.T) {

	c1 := circle{
		radius: 6.0,
	}
	got := c1.area()

	expected := 113.1

	equal := withTolerance(got, expected)

	if equal != true {
		t.Errorf("got %f, expected %f", got, expected)
	}

}

func TestCirclePerimeter(t *testing.T) {
	
	c1 := circle{
		radius: 6.0,
	}
	got := c1.perimeter()
	
	expected := 37.69911

	equal := withTolerance(got, expected)

	if equal != true {
		t.Errorf("got %f, expected %f", got, expected)
	}

}