package main

import "fmt"

// FUNGSI YANG MEMANGGIL DIRINYA SENDIRI

func Faktorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Faktorial(n-1)
}

func main() {
	fmt.Println(Faktorial(7))
}
