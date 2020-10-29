package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number *int
	var name *string

	var angka int
	var nama string

	fmt.Println(number, name)
	fmt.Println(reflect.TypeOf(number), reflect.TypeOf(name))
	fmt.Println(angka)
	fmt.Print(nama, "\n\n")

	angka = 5
	nama = "Muhammad Mirakel"

	fmt.Println(angka)
	fmt.Println(nama)

	number, name = &angka, &nama
	fmt.Println(number, name)
	fmt.Print(*number, " ", *name, "\n\n")

	angka = 7
	fmt.Println(*number)
	fmt.Print(angka, "\n\n")

	fmt.Println(number)
	fmt.Println(&angka)
}
