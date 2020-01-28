package main

import (
	"fmt"
)

func main() {

	v := func(nilai int) int {
		return nilai
	}

	result := checkNilai(v)
	fmt.Println(result)

	name, nilai := checkData("Rachmad Kurniawan", v)
	fmt.Println(name, nilai)
}

func checkNilai(nilai func(int) int) int {
	return nilai(10)
}

func checkData(name string, nilai func(int) int) (string, int) {
	return name, nilai(100)
 }