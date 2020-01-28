package main

import (
	"fmt"
)

// Profil
type (
	Education struct {
		SMA string
		SMP string
	}

	Profil struct {
		ID        int
		Name      string
		Age       int
		Address   string
		Education Education
	}
)

func main() {
	profil := Profil{
		ID:      1,
		Name:    "Rachmad Kurniawan",
		Age:     26,
		Address: "Kembangan, Jakarta Barat",
		Education: Education{
			SMA: "SMA N 1 Moutong",
			SMP: "SMP Al-Azhar Palu",
		},
	}

	fmt.Println(profil)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", profil.Name)
	fmt.Printf("Umur \t: %d\n", profil.Age)
	fmt.Printf("Alamat \t: %v\n", profil.Address)
	fmt.Printf("Sekolah \t: %v", profil.Education)
}