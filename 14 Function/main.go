package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// FUNGSI BIASA TANPA PARAMETER DAN TANPA NILAI KEMBALIAN
func Cetak_nama() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Printf("Halo, %s!\n", nama)
}

// FUNGSI TANPA PARAMETER DENGAN NLIAI KEMBALIAN
func Get_tanggal() string {
	return time.Now().Format("02-01-2006")
}

// FUNGSI BERPARAMETER TIPE DATA SAMA
func Luas_segitiga(alas, tinggi int) float32 {
	luas := float32(alas * tinggi / 2)
	return luas
}

// FUNGSI DENGAN PARAMETER BERTIPE DATA BEDA
func body_mass_index(tinggi float32, berat_badan int) float32 {
	bmi := float32(berat_badan) / tinggi * tinggi
	return bmi
}

func main() {
	Cetak_nama()
	Cetak_nama()

	fmt.Println(Get_tanggal())

	luas := Luas_segitiga(5, 10)
	fmt.Println(luas)

	bmi_var := body_mass_index(1.65, 70)
	fmt.Println(bmi_var)
}
