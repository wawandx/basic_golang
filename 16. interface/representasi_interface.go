package main

import (
	"fmt"
)

type Kendaraan interface {
	GetWarna()
}

type Mobil struct {
	Depan    string
	Belakang string
	Atas     string
	Dalam    string
}

func (m Mobil) GetWarna() {
	fmt.Println(m)
}

func CetakWarna(w Kendaraan) {
	fmt.Println(w)
}

func main() {
	var k Kendaraan
	mobil := Mobil{
		Depan:    "Biru",
		Belakang: "Biru",
		Atas:     "Hitam",
		Dalam:    "Putih",
	}

	k = mobil
	CetakWarna(k)
}