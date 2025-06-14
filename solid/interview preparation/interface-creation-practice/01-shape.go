//Shape Interface
//Create a Shape interface with methods Area() and Perimeter(). Implement this interface for Circle and Rectangle structs.

package interface_creation_practice

import (
	"fmt"
	"math"
)

type Shape interface {
	IArea() float64
	IPerimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) IArea() float64 {
	return r.Length * r.Width
}

func (r Rectangle) IPerimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c Circle) IArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) IPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	circleObject := Circle{
		Radius: 10.0,
	}
	areaOfCircle := circleObject.IArea()
	fmt.Println("Area of Circle : ", areaOfCircle)
}
