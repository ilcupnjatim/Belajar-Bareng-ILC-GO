package main

import "fmt"

func main() {
	// --------- Inisialisasi variabel array ---------

	scores := [3]int{100, 80, 90}

	fmt.Println(scores) // [100 80 90]

	nums := [...]int{100, 80, 90}

	fmt.Println(nums) // [100 80 90]
	fmt.Println()

	// --------- Mengakses suatu nilai elemen array ---------

	scores = [3]int{100, 80, 90}

	fmt.Println(scores[0]) // 100
	fmt.Println(scores[1]) // 80
	fmt.Println(scores[2]) // 90
	fmt.Println()

	// --------- Mengubah nilai elemen array ---------

	scores = [3]int{100, 80, 90}

	scores[0] = 95

	fmt.Println(scores) // [95 80 90]
	fmt.Println()

	// --------- Array multidimensi ---------

	numbers := [2][3]int{
		{10, 20, 30},
		{40, 50, 60},
	}

	fmt.Println(numbers) // [[10 20 30] [40 50 60]]

	fmt.Println(numbers[0]) // [10 20 30]

	fmt.Println(numbers[0][1]) // 20
}
