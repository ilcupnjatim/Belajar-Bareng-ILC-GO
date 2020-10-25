package main

import "fmt"

func main() {
	// --------- Operasi logika 2 buah boolean ---------

	fmt.Println(false && true) // false

	fmt.Println(false || true) // true

	fmt.Println(!true) // false

	// --------- Operasi logika 3 buah boolean ---------

	fmt.Println(false && true || true) // true

	// --------- Kombinasi operasi logika & perbandingan ---------

	const nilai = 80
	const absensi = 85

	const lulus = nilai >= 80 && absensi >= 90

	fmt.Println("Lulus:", lulus) // Lulus: false
}
