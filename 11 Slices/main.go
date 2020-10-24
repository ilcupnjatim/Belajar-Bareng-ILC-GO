package main

import "fmt"

func main() {
	// --------- Inisialisasi slice ---------

	friends := []string{"Udin", "Sohib", "Jafar"}

	fmt.Println(friends)      // [Udin Sohib Jafar]
	fmt.Println(friends[1])   // Sohib
	fmt.Println(cap(friends)) // 3
	fmt.Println()

	// --------- Membuat slice dari sebuah array ---------

	days := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	sliceDays := days[1:4]

	fmt.Println(sliceDays)      // [Selasa Rabu Kamis]
	fmt.Println(len(sliceDays)) // 3
	fmt.Println(cap(sliceDays)) // 6
	fmt.Println()

	// --------- Mengubah data slice ---------

	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Minggu]

	sliceDays[0] = "Seloso"

	fmt.Println(days) // [Senin Seloso Rabu Kamis Jumat Sabtu Minggu]
	fmt.Println()

	// --------- Menambahkan data ke dalam slice ---------

	days = []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	sliceDays = days[1:4]
	fmt.Println(sliceDays) // [Selasa Rabu Kamis]

	newSliceDays := append(sliceDays, "Liburr")
	fmt.Println(newSliceDays) // [Selasa Rabu Kamis Liburr]
	fmt.Println(days)         // [Senin Selasa Rabu Kamis Liburr Sabtu Minggu]

	fmt.Println()

	// --------------------------------------
	// Menambahkan data ke dalam slice melebihi kapasitasnya

	days = []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	sliceDays = days[4:]
	fmt.Println(sliceDays) // [Jumat Sabtu Minggu]

	newSliceDays = append(sliceDays, "Liburrr")
	fmt.Println(newSliceDays) // [Jumat Sabtu Minggu Liburrr]
	fmt.Println(days)         // [Senin Selasa Rabu Kamis Jumat Sabtu Minggu]

	fmt.Println()

	// --------- Membuat slice baru menggunakan make() ---------

	nums := make([]int, 2, 5)
	nums[0] = 10
	nums[1] = 20

	fmt.Println(nums) // [10 20]
	fmt.Println()

	// --------- Membuat slice baru menggunakan make() ---------

	sourceSlice := []int{10, 20, 30}
	fmt.Println(sourceSlice) // [10 20 30]

	destSlice := make([]int, 3, 3)
	copy(destSlice, sourceSlice)
	fmt.Println(destSlice) // [10 20 30]

	fmt.Println()

	// --------- Membuat multidimensinal slice ---------

	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(values)
}
