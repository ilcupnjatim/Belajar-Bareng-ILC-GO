package main

import "fmt"

func main() {
	// --------- Keyword if ---------

	const name = "udin"

	if name == "udin" {
		fmt.Println("Hello udin 😊")
	}

	// Output:
	// Hello udin 😊

	// ------------------------------

	const num = 1

	if num == 2 {
		fmt.Println("num bernilai 2")
	}

	// Output:
	//

	// ------------------------------

	const nilai = 80
	const absensi = 85

	if nilai >= 80 && absensi >= 75 {
		fmt.Println("Selamat, anda lulus!")
	}

	// Output:
	// Selamat, anda lulus!

	// --------- Keyword if-else ---------

	const lulus = false

	if lulus {
		fmt.Println("Anda lulus! 😊")
	} else {
		fmt.Println("Anda tidak lulus! 😢")
	}

	// Output:
	// Anda tidak lulus! 😢

	// --------- Keyword if, else if, else ---------

	const yourName = "udin"

	if yourName == "kacong" {
		fmt.Println("Hello Kacong")
	} else if yourName == "udin" {
		fmt.Println("Hello Udin")
	} else {
		fmt.Println("Ga kenal 😐")
	}

	// Output:
	// Hello Udin

	// --------- Keyword if dengan inisialisasi variabel ---------

	const myName = "Amir Muhammad Hakim"

	if length := len(myName); length > 10 {
		fmt.Printf("Nama dengan %d huruf terlalu panjang!", length)
	}

	// Output:
	// Nama dengan 19 huruf terlalu panjang!
}
