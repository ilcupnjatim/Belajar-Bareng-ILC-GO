package main

import "fmt"

type user struct {
	ID       string
	name     string
	email    string
	isActive bool
}

type group struct {
	name  string
	admin user
	users []user
}

func displayUser(u user) {
	fmt.Println("Nama:", u.name)
	fmt.Println("Email:", u.email)
	fmt.Println()
}

type human struct {
	name string
	age  int
}

type student struct {
	human
	grade int
}

func main() {
	// ---------- Basic Struct ----------

	bayu := user{
		ID:       "u1",
		name:     "Uzumaki Bayu",
		email:    "uzumaki.bayu@gmail.com",
		isActive: true,
	}

	fmt.Println(bayu)

	saburo := user{}
	saburo.ID = "u2"
	saburo.name = "Uzumaki Saburo"
	saburo.email = "uzumaki.saburo@outlook.com"
	saburo.isActive = false

	fmt.Println(saburo.name)
	fmt.Println(saburo.isActive)
	fmt.Println()

	// ---------- Objek Struct Bebagai Parameter Fungsi  ----------

	displayUser(saburo)

	// ---------- Nested Struct ----------

	binomo := user{
		ID:       "u3",
		name:     "Uzumaki binomo",
		email:    "uzumaki.binomo@gmail.com",
		isActive: true,
	}

	ilc := group{
		name:  "Informatic Learning Community",
		admin: binomo,
		users: []user{bayu, saburo},
	}

	fmt.Println("Nama grup:", ilc.name)
	fmt.Println("Nama admin:", ilc.admin.name)
	fmt.Println("Nama anggota 1:", ilc.users[0].name)
	fmt.Println()

	// ---------- Embedded struct ----------

	amir := student{}
	amir.name = "Amir Hakim"
	amir.age = 21
	amir.grade = 7

	fmt.Println("Name:", amir.name)
	fmt.Println("Age:", amir.age)
	fmt.Println("Grade:", amir.grade)
	fmt.Println()

	// ---------- Pointer Objek Struct ----------

	hakim := user{
		ID:       "u4",
		name:     "Hakim",
		email:    "hakim@gmail.com",
		isActive: true,
	}

	aku := &hakim

	fmt.Println("Nama hakim 1:", hakim.name)
	fmt.Println("Nama aku 1:", aku.name)

	hakim.name = "Amir Hakim"

	fmt.Println("Nama hakim 2:", hakim.name)
	fmt.Println("Nama aku 2:", aku.name)
}
