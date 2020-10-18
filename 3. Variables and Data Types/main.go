package main

import (
	"fmt"
	"reflect"
)

func main() {
	// DEKLARASI VARIABEL MENGGUNAKAN var
	var hewan, umur = "beruang", 21

	fmt.Println("nama hewan", hewan)
	fmt.Println(reflect.TypeOf(hewan))

	fmt.Printf("\numur saya %d\n", umur)
	fmt.Println(reflect.TypeOf(umur))

	// DEKLARASI VARIABEL TANPA var
	kendaraan := "mobil"
	fmt.Println("nama kendaraan", kendaraan)
	fmt.Println(reflect.TypeOf(kendaraan))

	kendaraan = "Kereta"

	fmt.Println(kendaraan)
}
