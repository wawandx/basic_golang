package main

import (
	"fmt"
)

const UMR int = 2000000

type Pekerja struct {
	Nama string
	Gaji int
}

func (p Pekerja) lihatPekerja() bool {
	fmt.Println("Nama\t :", p.Nama)
	fmt.Println("Gaji\t :", p.Gaji)
	if p.Gaji < UMR {
		return false
	} else {
		return true
	}
}

func main() {
	p1 := Pekerja{
		Nama: "Rachmad Kurniawan",
		Gaji: 6000000,
	}

	fmt.Println(p1.lihatPekerja())
}