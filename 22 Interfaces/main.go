package main

import "fmt"

type bangunDatar interface {
	hitungLuas() float64
}

type persegi struct {
	sisi float64
}

func (p persegi) hitungLuas() float64 {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang float64
	lebar   float64
}

func (pp persegiPanjang) hitungLuas() float64 {
	return pp.panjang * pp.lebar
}

func cetakLuas(p bangunDatar) {
	fmt.Printf("Luas: %f\n", p.hitungLuas())
}

func main() {
	// ---------- Basic Interfaces ----------

	var karpet bangunDatar = persegi{sisi: 4.0}
	fmt.Println("Luas:", karpet.hitungLuas())

	var papan bangunDatar = persegiPanjang{
		panjang: 5.0,
		lebar:   2.0,
	}
	fmt.Println("Luas:", papan.hitungLuas())
	fmt.Println()

	// ---------- Mengapa Perlu Interfaces? ----------

	cetakLuas(karpet)
	cetakLuas(papan)
	fmt.Println()

	// ---------- Embedded Interfaces ----------

	var bayu manusia = user{nama: "Uzumaki Bayu"}

	bayu.jalan()
	bayu.lihat()
	fmt.Println()

	// ---------- Interface Kosong ----------

	var mix interface{}

	mix = "Uzumaki Bayu"
	mix = 2
	mix = false

	fmt.Println("mix:", mix)

	saburo := map[string]interface{}{
		"name": "Uzumaki Saburo",
		"age":  21,
	}
	fmt.Println("saburo:", saburo)

	mix = 5
	result := mix.(int) * 2
	fmt.Println("Result:", result)
}

type punyaMata interface {
	lihat()
}

type punyaKaki interface {
	jalan()
}

type manusia interface {
	punyaMata
	punyaKaki
}

type user struct {
	nama string
}

func (u user) lihat() {
	fmt.Println(u.nama, "melihat sesuatu...")
}

func (u user) jalan() {
	fmt.Println(u.nama, "berjalan...")
}
