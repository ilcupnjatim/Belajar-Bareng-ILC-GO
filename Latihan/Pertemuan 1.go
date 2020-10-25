// SOLUSI DARI PEMATERI
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var angka_dari_user int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		fmt.Print("Masukan angka tebakan : ")
		fmt.Scanln(&angka_dari_user)

		if angka_dari_user == rand.Intn(10) {
			fmt.Println("Tebakan anda benar")
			break
		}
	}
}

// SOLUSI TERCEPAT - OLEH HUSNI T
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// inisiasi angka tebakan
	rand.Seed(time.Now().UnixNano())
	var cobaTebak = rand.Intn(10)
	var tebakan, kesempatan int

	kesempatan = 3
	for tebakan != cobaTebak {
		if kesempatan > 0 {
			fmt.Print("Coba tebak angka ( jumlah kesempatan ", kesempatan, " ) : ")
			fmt.Scanln(&tebakan)
		}

		if tebakan == cobaTebak {
			fmt.Print("Anda benar!")
			break
		}

		if kesempatan == 0 {
			fmt.Print("Ah sayang sekali~ kesempatan anda telah habis, jawaban yang benar adalah : ", cobaTebak)
			fmt.Print("better luck next time!")
			break
		}

		kesempatan--
	}
}


// SOLUSI TER OKE - OLEH FAISAL HANIF
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var jawaban = rand.Intn(10)
	var kesempatan, tebakan int

	fmt.Print("Masukkan jumlah kesempatan untuk menebak: ")
	fmt.Scanln(&kesempatan)

	for i := 0; i < kesempatan; i++ {
		fmt.Printf("Masukkan tebakan (%d dari %d): ", i+1, kesempatan)
		fmt.Scanln(&tebakan)

		if tebakan == jawaban {
			fmt.Println("Tebakan anda benar!")
			break
		}

	}

}
