package main

import (
	"fmt"
)

func main() {
	score := 10

	if score > 7 {
		fmt.Println("Anda Lulus dengan Nilai", score)
	} else {
		fmt.Println("Anda Tidak Lulus dengan Nilai", score)
	}

	nilai := 70
	switch nilai {
	case 50:
		fmt.Println("Nilai Anda C")
	case 70:
		fmt.Println("Nilai Anda B")
	case 100:
		fmt.Println("Nilai Anda A")
	default:
		fmt.Println("Tidak ada kategori")
	}
}