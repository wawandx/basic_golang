package main

import (
	"fmt"
)

// Profil
type (
	Profil struct {
		ID      int
		Name    string
		Age     int
		Address string
	}
)

func main() {
	profil := []Profil{
		{
			ID:      1,
			Name:    "Rachmad Kurniawan",
			Age:     26,
			Address: "Kembangan, Jakarta Barat",
		},
		{
			ID:      2,
			Name:    "Riski Anggita",
			Age:     19,
			Address: "Palu",
		},
	}

	for key, val := range profil {
		fmt.Println("Data Nama ke ", key, "Adalah : ", val.Name)
	}

}
