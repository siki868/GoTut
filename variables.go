package main

import "fmt"

/*

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte (alias for uint8)

rune -alias for int32 (represents a Unicode code point)

float32 float64

complex64 complex128

*/

// MAKSI Konstanta, ne moze := za deklarisanje
const MAKSI = 50

// Obicna
func add(x, y int) int {
	return x + y
}

// Swap f-ja
func swap(a, b string) (string, string) {
	return b, a
}

// Naked f-ja
func naked(sum int) (x, y int) {
	x = sum / 2
	y = x - 1 + sum
	return
}

func main0() {

	// Swap
	a, b := swap("beep", "boop")
	fmt.Println(a, b)

	// Naked
	fmt.Println(naked(23))

	// Var
	var c, python, java bool
	fmt.Println(c, python, java)

	// := dodeljivanje bez var-a, mozem da se koristi samo u telu f-je
	var i, j = 2, .1
	k := "aha"
	fmt.Println(i, j, k)

}
