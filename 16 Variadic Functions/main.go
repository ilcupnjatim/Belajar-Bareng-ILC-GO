package main

import (
	"fmt"
	"reflect"
)

// SIMPLE VARIADIC FUNCTION
func Himpunan_angka(angka ...int) {
	fmt.Println(angka)
	fmt.Println(reflect.TypeOf(angka))
}

// VARIADIC FUNCTION DENGAN PARAMETER LEBIH DARI 1
func Pembagi(pembagi int, angka_acak ...int) float32 {
	total := 0

	for _, num := range angka_acak {
		total += num
	}

	return float32(total / pembagi)
}

// func Nama_dan_angka(nama ...string, angka_acak ...int) {
// 	fmt.Println(nama)
// 	fmt.Println(angka)
// }

func main() {
	Himpunan_angka(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	hasil := Pembagi(5, 1, 2, 3, 4, 5)
	fmt.Println(hasil)

	//Nama_dan_angka("Jojo", "Naruto", "Ichigo", 12, 21, 666)
}

// VARIADIC FUNCTION DENGAN PARAMETER BEDA TIPE DATA DAN ELEMEN TAK TERBATAS?
// PAKE interface{}
