package main

import "fmt"

func Ubah_lewat_fungsi(awal *int, akhir int) {
	*awal = akhir
}

func main() {
	angka := 4
	fmt.Println("Awal :", angka)

	Ubah_lewat_fungsi(&angka, 7)
	fmt.Println("Akhir :", angka)
}
