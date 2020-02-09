package main

import "fmt"

/*
	recover
	panic
*/

// defer izvrsava ovu fju koja proverava da li je recover panic ili ne, ako je nil nije
func vrati() {
	if r := recover(); r != nil {
		fmt.Println("Izvukao se iz ", r)
	}
}

// Puca kad naidje na 4
func radi() {
	defer vrati()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			panic("Oh ne cetvorka, moram da crknem")
		}
	}
}

func main() {
	radi()
}
