package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var umur, npm int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Masukan umur anda : ")
	fmt.Scanln(&umur)
	fmt.Print("Masukan NPM anda : ")
	fmt.Scanln(&npm)

	fmt.Print("Masukan nama kampus anda : ")
	scanner.Scan()
	kampus := scanner.Text()

	fmt.Println() // BUAT NAMBAH BARIS KOSONG AJA

	fmt.Println("Nama saya adalah", nama)
	fmt.Println("Umur saya adalah", umur, "tahun")
	fmt.Println("NPM saya adalah", npm)
	fmt.Println("Kampus saya adalah", kampus)
}
