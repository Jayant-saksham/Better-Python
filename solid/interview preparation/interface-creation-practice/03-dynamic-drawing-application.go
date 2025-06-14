// Dynamic Drawing Application
// Define a Drawable interface with Draw() method. Implement it for different types like Circle, Line, and TextBox. Then write a function that takes a slice of Drawable and calls Draw() on each.

package interface_creation_practice

import "fmt"

type Drawable interface {
	IDraw()
}

type CircleDraw struct{}
type RectangleDraw struct{}
type LineDraw struct{}

func (c CircleDraw) IDraw() {
	fmt.Println("CircleDraw")
}
func (r RectangleDraw) IDraw() {
	fmt.Println("RectangleDraw")
}
func (l LineDraw) IDraw() {
	fmt.Println("LineDraw")
}

func CallDrawForEach(drawable []Drawable) {
	for _, d := range drawable {
		d.IDraw()
	}
}

func main() {
	circleDrawable := CircleDraw{}
	rectDrawable := RectangleDraw{}
	lineDrawable := LineDraw{}

	listOfDrawable := []Drawable{circleDrawable, rectDrawable, lineDrawable}

	CallDrawForEach(listOfDrawable)
}
