package main

import (
	"fmt"
)

type Pekerja struct {
	Nama string
	Gaji int
}

func (p Pekerja) lihatPekerja() {
	fmt.Println("Nama\t :", p.Nama)
	fmt.Println("Gaji\t :", p.Gaji)
}

func main() {
	p1 := Pekerja{
		Nama: "Rachmad Kurniawan",
		Gaji: 6000000,
	}

	p1.lihatPekerja()
}