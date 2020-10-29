package main

import (
	"bufio"
	"fmt"
	"os"
)

// NILAI KEMBALIAN TANPA MENYEBUTKAN VARIABEL DI AKHIR FUNGSI
func Get_Nama() (nama string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	nama = scanner.Text()

	return
}

// NILAI KEMBALIAN GANDA
func Ganda(nama string, umur int) (string, int) {
	return nama, umur * 2
}

func main() {
	nama_ku := Get_Nama()

	a, b := Ganda(nama_ku, 21)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	_, c := Ganda(nama_ku, 22)
	fmt.Println(c)
}
