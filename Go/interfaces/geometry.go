package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return (r.width + r.height) * 2
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	dim := fmt.Sprintf("%3.1f", g)
    fmt.Printf("%-10s Area:%3.3f Perim:%3.3f\n",dim, g.area(),g.perimeter())
}
func Geometry() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
	measure(r)
    measure(c)
}