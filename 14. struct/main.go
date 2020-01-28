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
	var profil Profil

	profil.ID = 1
	profil.Name = "Rachmad Kurniawan"
	profil.Age = 26
	profil.Address = "Kembangan, Jakarta Barat"

	fmt.Println(profil)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", profil.Name)
	fmt.Printf("Umur \t: %d\n", profil.Age)
	fmt.Printf("Alamat \t: %v\n", profil.Address)
	fmt.Println("===========================")
}
