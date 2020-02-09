package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Pravi se novi reader
	r := strings.NewReader("osamosam po osam")

	// Inicijalizuje se niz od 8 bajtova
	b := make([]byte, 8)

	for {
		// Reader cita 8 po 8 bajtova i ucitava ih u b niz
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
