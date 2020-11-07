package main

import (
	"23_packages/calculations"
	"23_packages/models"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// ---------- Akses Properties dari Package yang SAMA ----------

	// sentence := sayHello()

	// fmt.Println(sentence)
	// fmt.Println()

	// ---------- Akses Properties dari Package yang BERBEDA ----------

	result := calculations.Add(2, 3)

	fmt.Println("Result:", result)
	fmt.Println()

	// ---------- Akses Struct dari Package yang BERBEDA ----------

	bayu := models.User{
		ID:       "u1",
		Name:     "Uzumaki Bayu",
		Email:    "uzumaki.bayu@gmail.com",
		IsActive: true,
	}

	fmt.Println("Bayu:", bayu)
	fmt.Println()

	// ---------- Akses Property dari Package Library ----------

	fmt.Println("UUID:", uuid.New().String())
}
