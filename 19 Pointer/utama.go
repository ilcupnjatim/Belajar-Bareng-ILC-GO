package main

import "fmt"

func Ubah_lewat_pointer(awal *int, akhir int) {
	*awal = akhir
}

func main() {
	angka := 4
	fmt.Println("Awal :", angka)

	Ubah_lewat_pointer(&angka, 7)
	fmt.Println("Akhir :", angka)
}
