package main

import (
	"fmt"
	"math"
	"time"
)

// Stack defer
func deferStacked() {
	fmt.Println("Pocinjem loop")

	for i := 0; i < 10; i += 2 {
		defer fmt.Println(i)
	}
	fmt.Println("Kraj")
}

// Primer defera, posle defer-a ide funkcija
func deferTestic() {
	defer fmt.Println("Ja se izvrsavam posle")
	for i := 0; i < 10; i++ {
		// Nesto radim
		fmt.Println(i * 21)
	}

	fmt.Println("Kraj fj-e")
}

// ovakav switch je dobar za duge if-ove
func swLong() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Jutric")
	case t.Hour() < 17:
		fmt.Println("Podne")
	default:
		fmt.Println("Veche")
	}
}

// Primere switcha
func sw() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Danas!")
	case today + 1:
		fmt.Println("Sutra.")
	case today + 2:
		fmt.Println("Za 2 dana")
	default:
		fmt.Println("Tugy :(")
	}
}

func main() {
	// ez for
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while je for
	sum = 1
	for sum < 50 {
		sum += sum
	}
	fmt.Println(sum)

	// klasican if
	if 10 < 50 {
		fmt.Println("False")
	}

	// Moze da ima pocetnu dekleraciju kao i for i da ima return
	if k := add(20, 30); k < sum {
		fmt.Println(k, sum, math.Abs(float64(k-sum)))
	}

	// switch
	sw()
	swLong()

	// defer
	deferTestic()
	deferStacked()
}
