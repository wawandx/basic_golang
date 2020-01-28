package main

import (
	"fmt"
)

func ubah(data *int) {
	*data = 100
}

func main() {

	var a int = 10

	var b *int = &a

	ubah(b)

	fmt.Println("Alamat Memori variable b adalah", b)
	fmt.Println("Nilai dari variable b adalah", *b)
}