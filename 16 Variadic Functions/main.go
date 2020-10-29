package main

import (
	"fmt"
	"reflect"
)

// SIMPLE VARIADIC FUNCTION
func Tambah(nums ...int) {
	fmt.Println(nums)
	fmt.Println(reflect.TypeOf(nums))
}

// VARIADIC FUNCTION DENGAN PARAMETER LEBIH DARI 1
func Pembagi(pembagi int, acak ...int) float32 {
	total := 0

	for _, num := range acak {
		total += num
	}

	return float32(total / pembagi)
}

func main() {
	Tambah(1, 2, 3)

	angka := []int{1, 2, 3, 4}
	Tambah(angka...)

	fmt.Println(Pembagi(5, 6, 7, 8, 9))
}

// VARIADIC FUNCTION DENGAN PARAMETER BEDA TIPE DATA DAN ELEMEN TAK TERBATAS?
// PAKE interface{}
