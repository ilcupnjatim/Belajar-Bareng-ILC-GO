package main

import (
	"fmt"
	"runtime"
)

func cetak(n int, message string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, message)
	}
}

func main() {
	// ---------- Basic Goroutines  ----------

	// runtime.GOMAXPROCS(4)

	// go cetak(5, "Hello")
	// go cetak(5, "Oii")
	// cetak(5, "Hai")

	// var input string
	// fmt.Scanln(&input)

	// ---------- Goroutines & Channels  ----------

	runtime.GOMAXPROCS(4)

	msgChannel := make(chan string)

	sayHelloTo := func(name string) {
		data := fmt.Sprintln("Hello", name)
		msgChannel <- data
	}

	go sayHelloTo("Bayu")
	go sayHelloTo("Saburo")
	go sayHelloTo("Binomo")

	msg1 := <-msgChannel
	fmt.Println("msg1:", msg1)

	msg2 := <-msgChannel
	fmt.Println("msg2:", msg2)

	msg3 := <-msgChannel
	fmt.Println("msg3:", msg3)
}
