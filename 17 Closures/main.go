package main

import (
	"fmt"
	"reflect"
)

// Closure adalah fungsi yang bisa langsung disimpan
// di sebuah variabel pada saat inisialisasi'nya

func main() {
	kali_dua := func(angka int) (hasil float32) {
		hasil = float32(angka * 2)
		return
	}

	fmt.Println(kali_dua(5))
	fmt.Println(reflect.TypeOf(kali_dua(5)))

	// Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	// atau bahkan membuat fungsi yang mengembalikan fungsi.

	// Closure merupakan anonymous function atau fungsi tanpa nama. Biasa
	// dimanfaatkan untuk membungkus suatu proses yang hanya dipakai pada blok tertentu saja.

	// Immediately-Invoked Function Expression
	hewan_kurban := []string{"Sapi", "Kambing", "Domba", "Unta"}

	beli_satu := func(indeks int) string {
		return hewan_kurban[indeks]
	}(2)

	fmt.Println(beli_satu)

	kubik := Pangkat_empat(2)
	fmt.Println(kubik())
}

// FUNGSI YANG MENGEMBALIKAN FUNGSI
func Pangkat_empat(bilangan int) func() int {
	bilangan = bilangan * bilangan

	return func() int {
		return bilangan * bilangan
	}
}
