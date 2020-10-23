package main

import "fmt"

func main() {
	// --------- Membandingkan 2 buah string ---------

	const name1 = "udin"
	const name2 = "udin"

	const isEqual = name1 == name2

	fmt.Println(isEqual) // true
	fmt.Println()

	// --------- Membandingkan 2 buah integer ---------

	const num1 = 2
	const num2 = 3

	fmt.Println(num1 > num2)  // false
	fmt.Println(num1 < num2)  // true
	fmt.Println(num1 == num2) // false
	fmt.Println(num1 != num2) // true
	fmt.Println()

	// --------- Membandingkan 2 buah boolean ---------

	const bool1 = true
	const bool2 = false

	fmt.Println(bool1 == bool2) // false
	fmt.Println(bool1 != bool2) // true
	fmt.Println()
}
