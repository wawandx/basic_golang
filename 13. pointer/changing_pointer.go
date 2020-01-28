package main

import (
	"fmt"
)

func ubah(data *int) {
	*data = 100
}

func ubah2(data int) *int {
	data = 20
	return &data
}

func main() {

	var a int = 10
	var b *int = &a
	var c int = 15

	ubah(b)
	datac := ubah2(c)

	fmt.Println("Alamat Memori variable b adalah", b)
	fmt.Println("Nilai dari variable b adalah", *b)

	fmt.Println("Alamat Memori variable c adalah", datac)
	fmt.Println("Nilai dari variable c adalah", *datac)
}
