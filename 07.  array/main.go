package main

import (
	"fmt"
)

func main() {
	var name [5]string
	name[0] = "Rachmad"
	name[1] = "Kurniawan"
	name[3] = "Riski"
	name[4] = "Anggita"

	for index := 0; index < len(name); index++ {
		fmt.Println("Index ke ", index, "=>", name[index])
	}
}