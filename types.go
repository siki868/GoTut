package main

import (
	"fmt"
)

/*
	Closure functions
	Maps
	Arrays
	Slicing
	Make
*/

// Funkcije mogu da se prosledjuju u funkcije kao promenljive
func doFunc(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

// closure f-ja, vraca funkciju koja vraca int
func sabirac() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// closure f-ja za fibonacci
func fib() func() int {
	pn, pk := 0, 1

	return func() (p int) {
		p = pn
		pn, pk = pk, pn+pk
		return p
	}
}

func main4() {

	// Klasicno slajsovanje iz python-a
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	var s []int = primes[3:]
	fmt.Println(s)

	// Niz struktura bez deklarisanja duzine
	st := []struct {
		i   int
		str string
	}{
		{1, "Jedan"},
		{2, "Dva"},
		{3, "Tri"},
		{4, "Cetiri"},
	}
	fmt.Println(st)

	// Dinamicko pravljenje slice-a
	// len - duzina slice-a, cap - kapacitet slice-a
	a := make([]int, 5)
	fmt.Printf("%s len=%d cap=%d %v\n", "a", len(a), cap(a), a)

	// Appendovanje slice-ova
	var niz []int
	niz = append(niz, 0)
	fmt.Println(niz)
	niz = append(niz, 1, 2)
	fmt.Println(niz)

	// Loopovanje kroz slice sa range-om, uvek se vracaju 2 vrednosti, index i value
	a = []int{1, 2, 4, 34, 12, 100}
	for i, v := range a {
		fmt.Printf("%d. %d\n", i, v)
	}

	// Za skipovanje promenjlive isto kao i u py _
	a = make([]int, 10)
	for i := range a {
		a[i] = 1 << uint(i) // 2^i
	}
	for _, value := range a {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// mape - key:value, u ovom slucaju str:int, inicijalizuje se sa make
	var m map[string]int
	m = make(map[string]int)
	m["Prvi"] = 10
	m["Drugi"] = 230
	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
	// Provera da l' postoji key:value u mapi
	elem, ok := m["Nema me"]
	if !ok {
		fmt.Printf("Ne postoji, elem je %v\n", elem)
	}
	// Brisanje
	fmt.Printf("Brisem elem 'Prvi'\n")
	delete(m, "Prvi")
	fmt.Println(m)

	// fje
	saberi := func(a, b int) int {
		return a + b
	}
	fmt.Println(doFunc(saberi, 10, 20))

	// closure f-je, prate se 2 razlicite sume na razl nacine
	s1, s2 := sabirac(), sabirac()
	for i := 0; i < 5; i++ {
		fmt.Println(s1(i), s2(-i+3*i))
	}

	// Primer fib closure fj-e
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Println()

}
