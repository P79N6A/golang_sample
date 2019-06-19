package hello

import "fmt"

//Vertex {x, y} type
type Vertex struct {
	X, Y int
}

//String for Vertex
func (v *Vertex) String() string {
	return fmt.Sprintf("(x=%v, y=%v)", v.X, v.Y)
}
