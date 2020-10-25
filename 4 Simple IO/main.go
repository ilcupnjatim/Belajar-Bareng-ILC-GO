package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	//var umur, npm int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("inputan : ", line)

	fmt.Print("Masukan nama anda :")
	fmt.Scanln(&nama)
	/*
		fmt.Print("Masukan umur anda :")
		fmt.Scanln(&umur)
		fmt.Print("Masukan NPM anda :")
		fmt.Scanln(&npm)
		fmt.Print("Masukan nama kampus anda :")
		fmt.Scanln(&kampus)
	*/

	fmt.Println()
	fmt.Println("Nama saya adalah", nama)
	/*
		fmt.Println("Umur saya adalah", umur, "tahun")
		fmt.Println("NPM saya adalah", npm)
		fmt.Println("Kampus saya adalah", kampus)
	*/
}
