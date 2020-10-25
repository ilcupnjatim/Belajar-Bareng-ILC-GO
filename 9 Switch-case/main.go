package main

import "fmt"

func main() {
	// --------- Basic switch-case ---------

	const myAge = 21

	switch myAge {
	case 20:
		fmt.Println("Umur saya dua puluh tahun")
	case 21:
		fmt.Println("Umur saya dua puluh satu tahun")
		fmt.Println("Hari ini saya tambah umur 🥳")
	default:
		fmt.Println("Umur saya tidak diketahui")
	}

	// Output:
	// Umur anda dua puluh satu tahun
	// Hari ini saya tambah umur 🥳

	// --------- Sebuah case dengan banyak kondisi ---------

	const yourAge = 22

	switch yourAge {
	case 20:
		fmt.Println("Umur kamu 21 tahun")
	case 21, 22, 23:
		fmt.Println("Umur kamu 21 / 22 / 23 tahun")
	default:
		fmt.Println("Umur kamu tidak diketahui")
	}

	// Output:
	// Umur anda 21 / 22 / 23 tahun

	// --------- Switch-case dengan gaya if statements ---------

	const name = "Amir Muhammad Hakim"

	const length = len(name)

	switch {
	case length > 10:
		fmt.Printf("Nama dengan %d huruf terlalu panjang!\n", length)
	case length > 5:
		fmt.Printf("Nama dengan %d huruf masih terlalu panjang!\n", length)
	}

	// Output:
	// Nama dengan 19 huruf masih terlalu panjang!

	// --------- Switch-case dengan inisialisasi variabel ---------

	const firstName = "Amir"

	switch length := len(firstName); length > 10 {
	case true:
		fmt.Printf("Nama dengan %d huruf terlalu panjang!\n", length)
	case false:
		fmt.Println("Nama sudah pas 👍")
	}

	// Output:
	// Nama sudah pas 👍
}
