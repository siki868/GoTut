package main

/*
	Methods
	Inerfaces
*/

import (
	"fmt"
	"math"
	"strings"
)

// Inter : definisanje interface-a
type Inter interface {
	pisi()
}

// Vertex : Go nema klase, metode se definisu preko f-ja
type Vertex struct {
	X, Y float64
}

// Metoda se pravi tako sto se izmedju ime f-je i func keyworda stavi RECEIVER, u ovom slucaju abs metoda ima receiver tipa Vertex sa imenom v
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Metoda moze da prima i pokazivac na strukturu
func (v *Vertex) skaliraj(s float64) {
	v.X *= s
	v.Y *= s
}

// Tip vertex implementira interface Inter ali ne treba implicitno da se navodi
func (v Vertex) pisi() {
	fmt.Printf("(%.2f %.2f)\n", v.X, v.Y)
}

// Interface je univerzalan i sastoji se od (value, type) para
func interfaceTest(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Eksponent: ", math.Exp2(float64(v)))
		break
	case string:
		fmt.Println("Velika: ", strings.ToUpper(v))
		break
	default:
		fmt.Printf("Ovo je tipa %T\n", v)
		break
	}
}

// toString metoda
func (v Vertex) String() string {
	return fmt.Sprintf("Koordinate: (%.2f, %.2f)", v.X, v.Y)
}

func main1() {
	// Koriscenje metode
	v := Vertex{3, 4}
	fmt.Println(v.abs())
	scale := 3.8123
	v.skaliraj(scale)
	fmt.Printf("Sklairano za %v: (%.2f %.2f)\n", scale, v.X, v.Y)
	fmt.Println(v)
	// Preko interface-a
	var i Inter = Vertex{1.2, 3.4412}
	i.pisi()

	// Testiranje intefejsa
	interfaceTest(15)
	interfaceTest("mala slova")
	interfaceTest(2.3)

}
