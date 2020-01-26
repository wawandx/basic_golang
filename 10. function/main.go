package main

import (
	"fmt"
)

func cetak() {
	fmt.Println("Belajar Fungsi di Golang")
}

func tambah(a int, b int) int {
	c := a + b
	return c
 }

func main() {
	cetak()

	a := 10
	b := 20
	
	hasil := tambah(a, b)
	
	fmt.Println("Hasil Penjumlahan Antara", a, "+", b, "Adalah", hasil)
}