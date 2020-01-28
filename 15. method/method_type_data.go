package main

import (
	"fmt"
)
type val int

func (i val) jumlah(a int) val {
	s := i * val(a)
	return s
}

func main() {
	numb := val(10)
	fmt.Println(numb.jumlah(2))
}