package main

import (
	"fmt"
)

func main() {
	for index := 0; index < 10; index++ {
		fmt.Println("Looping in ", index)
	}

	var index = 0
	for index < 4 {
		fmt.Println("number of", index)
		index++
		if index == 2 {
			break
		}
	}

	data := []string{"Rachmad", "Kurniawan", "Anggita"}

	for index, v := range data {
		fmt.Println("looping in ", index, v)
	}
}
