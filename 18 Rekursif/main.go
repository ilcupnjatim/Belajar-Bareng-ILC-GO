package main

import "fmt"

// FUNGSI YANG MEMANGGIL DIRINYA SENDIRI
func Faktorial(angka int) int {
	if angka == 0 {
		return 1
	}

	return angka * Faktorial(angka-1)
}

// CONTOH DUA
func Fibonanci(angka int) int {
	if angka == 0 {
		return 0
	} else if angka == 1 {
		return 1
	} else {
		return Fibonanci(angka-1) + Fibonanci(angka-2)
	}
}

func main() {
	fmt.Println(Faktorial(5))

	for i := 0; i < 5; i++ {
		fmt.Println(Fibonanci(i))
	}
}
