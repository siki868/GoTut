package main

/*
	Methods
*/

import (
	"fmt"
	"math"
)

// Vertex : Go nema klase, metode se definisu preko f-ja
type Vertex struct {
	X, Y float64
}

// Metoda se pravi tako sto se izmedju ime f-je i func keyworda stavi RECEIVER, u ovom slucaju abs metoda ima receiver tipa Vertex sa imenom v
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Metoda moze da prima i pokazivac na tip
func (v *Vertex) skaliraj(s float64) {
	v.X *= s
	v.Y *= s
}

func main() {
	// Koriscenje metode
	v := Vertex{3, 4}
	fmt.Println(v.abs())
	scale := 3.8123
	v.skaliraj(scale)
	fmt.Printf("Sklairano za %v: (%.2f %.2f)\n", scale, v.X, v.Y)
}
