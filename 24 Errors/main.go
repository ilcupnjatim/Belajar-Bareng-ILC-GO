package main

import (
	"errors"
	"fmt"
)

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Terjadi pembagian dengan nol")
	}
	return a / b, nil
}

func main() {
	// ---------- Package "errors" & Error handling  ----------

	result, err := bagi(8, 2)
	if err != nil {
		fmt.Println("Terjadi error:", err.Error())
	} else {
		fmt.Println("Hasil 8 / 2 =", result)
	}

	result, err = bagi(8, 0)
	if err != nil {
		fmt.Println("Terjadi error:", err.Error())
	} else {
		fmt.Println("Hasil 8 / 0 =", result)
	}
}
