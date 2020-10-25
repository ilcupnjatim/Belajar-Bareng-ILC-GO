package main

import "fmt"

func main() {
	// ---------Contoh: iterasi array menggunakan for-loop ---------

	fruits := [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// Output:
	// apple
	// grape
	// banana
	// melon

	fmt.Println()

	// --------- Iterasi array menggunakan for-range ---------

	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// Output:
	// 0 apple
	// 1 grape
	// 2 banana
	// 3 melon

	fmt.Println()

	// --------- Iterasi slice menggunakan for-range ---------

	numbers := []int{20, 95, 100, 56}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Output:
	// 0 20
	// 1 95
	// 2 100
	// 3 56

	fmt.Println()

	// --------- Iterasi map menggunakan for-range ---------

	stockBuah := map[string]int{
		"apel":   21,
		"mangga": 22,
		"jeruk":  31,
	}

	for key, value := range stockBuah {
		fmt.Println(key, value)
	}

	// Output:
	// apel 21
	// mangga 22
	// jeruk 31
}
