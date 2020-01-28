package main

import (
	"fmt"
	"strings"
)

func main() {
	var mp interface{}

	//tipe slice
	mp = []string{"Satu", "Dua"}

	var toSlice = strings.Join(mp.([]string), ", ")
	fmt.Println(toSlice)
}