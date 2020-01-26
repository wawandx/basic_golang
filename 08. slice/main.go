package main

import (
	"fmt"
)

func main() {
	var mahasiswa []string
	mahasiswa = []string{"Rachmad Kurniawan", "Riski Anggita"}

	fmt.Println("Data Mahasiswa => ", mahasiswa)
	fmt.Println("Data Mahasiswa ke 0 sampai 1 => ", mahasiswa[0:1])
	fmt.Println("Data Mahasiswa ke 0 sampai 2 => ", mahasiswa[0:2])
	fmt.Println("Jumlah Mahasiswanya adalah => ", len((mahasiswa)))

	biologi := []string{"Rachmad Kurniawan", "Riski Anggita"}
	fisika := make([]string, len(biologi))
 
	copy(fisika, biologi)
	fmt.Println("Mahasiswa di Mata kuliah fisika", fisika)

	mahasiswaLama := []string{"Rachmad Kurniawan", "Riski Anggita"}
	mahasiswaBaru := "Rocky Rotikan"
	mahasiswaSlice := append(mahasiswaLama, mahasiswaBaru)
	fmt.Println(mahasiswaSlice)
}