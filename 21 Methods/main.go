package main

import "fmt"

type user struct {
	ID       string
	name     string
	email    string
	isActive bool
}

func (u user) display() {
	fmt.Println("Nama:", u.name)
	fmt.Println("Email:", u.email)
	fmt.Println()
}

func (u *user) updateName(newName string) {
	u.name = newName
}

func main() {
	// ---------- Struct Methods ----------

	bayu := user{
		ID:       "u1",
		name:     "Uzumaki Bayu",
		email:    "uzumaki.bayu@gmail.com",
		isActive: true,
	}

	bayu.display()

	binomo := user{
		ID:       "u2",
		name:     "Uzumaki Binomo",
		email:    "uzumaki.binomo@gmail.com",
		isActive: false,
	}

	binomo.display()

	// ---------- Update Property Struct Melalui Method ----------

	fmt.Println("Before:", binomo.name)

	binomo.updateName("Binomo")

	fmt.Println("After:", binomo.name)
	fmt.Println()
}
