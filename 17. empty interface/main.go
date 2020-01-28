package main

import (
	"fmt"
)

func cetakData(i interface{}) {
	fmt.Println(i)
}

func main() {
	// contoh tipe data string
	st := "Halo, belajar golang itu mudah!!!"
	cetakData(st)

	//tipe slice
	mp := []int{1, 2, 3, 4, 5}
	cetakData(mp)
}