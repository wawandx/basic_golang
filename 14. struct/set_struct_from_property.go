package main

import (
	"fmt"
)

// Profil
type Profil struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {

	profil := Profil{
		ID:      1,
		Name:    "Rachmad Kurniawan",
		Age:     26,
		Address: "Kembangan, Jakarta Barat",
	}

	fmt.Println(profil)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", profil.Name)
	fmt.Printf("Umur \t: %d\n", profil.Age)
	fmt.Printf("Alamat \t: %v", profil.Address)
}