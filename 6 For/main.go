package main

import "fmt"

func main() {
	// SEBAGAIMANA C
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	// while STYLE
	j := 0

	for j < 5 {
		fmt.Println(j)
		j = j + 1
	}

	fmt.Println("\nBREAK")
	// BREAK
	for k := 0; k < 10; k++ {

		if k == 5 {
			break
		}

		fmt.Println(k)
	}

	fmt.Print("\nCONTINUE\n")
	// CONTINUE
	for l := 0; l < 10; l++ {

		if l%2 == 0 {
			continue
		}

		fmt.Println(l)
	}
}
