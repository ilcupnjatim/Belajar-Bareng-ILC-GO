package main

import (
	"bufio"
	"fmt"
	"os"
)

// NILAI KEMBALIAN TANPA MENYEBUTKAN VARIABEL DI AKHIR FUNGSI
func Get_nama() (nama string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	nama = scanner.Text()

	return
}

// NILAI KEMBALIAN GANDA
func Ganda(nama string, umur int) (string, int) {
	return nama, umur
}

func main() {
	var umur_gue int

	fmt.Print("Masukan umur anda : ")
	fmt.Scanln(&umur_gue)

	nama, umur := Ganda(Get_nama(), umur_gue)

	fmt.Println(nama)
	fmt.Println(umur)

	_, umur_saya := Ganda(Get_nama(), umur_gue)
	fmt.Println(umur_saya)

}
