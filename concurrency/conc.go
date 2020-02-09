package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Goroutines
	Channels (kanali preko kojih se prosledjuju podaci)
	Buffered channels
	Select
	WaitGroup
*/
var wg sync.WaitGroup

// test fun za ispisivanje necega sa sleep-om
func pljuni(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

// fun za ispisivanje sa sleepom uz pomoc wait grupe
func pljuniWg(s string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

// test fun gde pokrecemo pljuni preko goroutine i normalno
func goTest() {
	// nova goroutine-a se pravi sa go, prakticno thread novi, goroutine imaju isti adresni prostor
	go pljuni("Go")
	pljuni("Ne go")
}

// go sa WaitGroupom koja omogucava sinhronizaciju groutine
func goTestWg() {
	wg.Add(1)
	go pljuniWg("Prvo")
	wg.Add(1)
	go pljuniWg("Drugo")
	wg.Wait()
}

// test fun gde sumu prosledjenog slice-a saljemo na proslednjeni kanal
func sumica(arr []int, c chan int) {
	suma := 0
	for _, v := range arr {
		suma += v
	}
	// sumu posalji na kanal c
	c <- suma
}

// testiranje kanala preko sume polovina niza
func chanTest() {
	arr := []int{1, 1, 1, 5, 5, 5}

	// inicijalizuje kanal za prosledjivanje intova
	c := make(chan int)

	// nadji sumu prve polovine niza
	go sumica(arr[:len(arr)/2], c)
	// nadju sumu druge polovine niza
	go sumica(arr[len(arr)/2:], c)

	// uzmi podatke sa kanala c
	// x ce biti suma drugog dela a y prvog
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

// Channelu moze da se da velicina buffera
func bufferecChann() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	fmt.Println(<-ch)
	defer fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// test fun za slanje svake sekvence fibonacija na kanal
func fibonacciChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// koriscenje prosle metode
func koristiFib() {
	c := make(chan int, 30)
	go fibonacciChan(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// select blok je blokiran sve dok ne moze da izvrsi neki od njegovih case-ova
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// koriscenje prosle fun
func koristiFibSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacciSelect(c, quit)

}

func main() {
	goTestWg()
}
