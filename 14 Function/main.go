package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// FUNGSI BIASA TANPA PARAMETER DAN TANPA NILAI KEMBALIAN
func Cetak_Nama() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Printf("Halo, %s!\n", nama)
}

// FUNGSI TANPA PARAMETER DENGAN NLIAI KEMBALIAN
func Get_Tanggal() string {
	return time.Now().Format("01-02-2006")
}

// FUNGSI BERPARAMETER TIPE DATA SAMA
func Luas_Segitia(alas, tinggi int) float32 {
	return float32(alas * tinggi / 2) // float32(alas) * float32(tinggi) / float32(2)
}

// FUNGSI DENGAN PARAMETER BERTIPE DATA BEDA
func body_mass_index(tinggi float32, berat_badan int) float32 {
	return float32(berat_badan) / (tinggi * tinggi)
}

func main() {
	Cetak_Nama()

	tgl_skrng := Get_Tanggal()
	fmt.Println(tgl_skrng)

	l_segi3 := Luas_Segitia(10, 5)
	fmt.Println(l_segi3)

	bmi := body_mass_index(1.65, 75)
	fmt.Println(bmi)
}
