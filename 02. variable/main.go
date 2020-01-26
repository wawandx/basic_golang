package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name, address, workplace string
	var age int
	const endPoint string = "https://www.ralali.com/api"
	
	name = "Rachmad Kurniawan"
	address = "Kembangan, Jakarta Barat"
	workplace = "Ralali.com"
	age = 26
	height := 178
	weight := 70
	fmt.Println("=== My Profile ===")
	fmt.Println("Name : " + name)
	fmt.Println("Address : " + address)
	fmt.Println("Workplace : " + workplace)
	fmt.Println("Age : ", age)
	fmt.Println("Height : ", height)
	fmt.Println("Weight : ", weight)
	fmt.Println("")
	fmt.Println("=== Constanta ===")
	fmt.Println("End Point %s", endPoint)

	var numberString = "794"
	var toInteger, err = strconv.Atoi(numberString)
 
	if err == nil {
		fmt.Println(toInteger)
	} else {
		fmt.Println(err)	
	}
 }