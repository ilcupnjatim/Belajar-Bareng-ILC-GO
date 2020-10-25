package main

import "fmt"

func main() {
	// --------- Membuat map ---------

	stockBuah := map[string]int{
		"apel":   21,
		"mangga": 22,
		"jeruk":  31,
	}

	fmt.Println(stockBuah) // map[apel:21 jeruk:31 mangga:22]
	fmt.Println()

	// --------- Mengakses elemen map & Mengubah elemen map ---------

	stockBuah = map[string]int{
		"apel":   21,
		"mangga": 22,
		"jeruk":  31,
	}

	fmt.Println(stockBuah["apel"])  // 21
	fmt.Println(stockBuah["jeruk"]) // 31

	stockBuah["apel"] = 50

	fmt.Println(stockBuah["apel"]) // 50
	fmt.Println()

	// --------- Menghapus elemen map ---------

	stockBuah = map[string]int{
		"apel":   21,
		"mangga": 22,
		"jeruk":  31,
	}

	delete(stockBuah, "jeruk")

	fmt.Println(stockBuah) // map[apel:21 mangga:22]
	fmt.Println()

	// --------- Membuat map dengan fungsi make()  ---------

	user := make(map[string]string)

	user["name"] = "udin"
	user["gender"] = "pria"

	fmt.Println(user) // map[gender:pria name:udin]
}
