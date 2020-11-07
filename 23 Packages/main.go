package main

import (
	"23_packages/calculations"
	"fmt"
)

func main() {
	// ---------- Akses Sesuatu dari Package yang SAMA ----------

	// sentence := sayHello()

	// fmt.Println(sentence)
	// fmt.Println()

	// ---------- Akses Sesuatu dari Package yang BEDA ----------

	result := calculations.Add(2, 3)

	fmt.Println("Result:", result)
	fmt.Println()
}
