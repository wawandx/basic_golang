package main

import (
	"fmt"
)

func main() {
	number1 := 10
	number2 := 5
	
	/**
		* OPERATOR ARITMATIKA
	*/
	//Operator plus
	value := number1 + number2
	fmt.Println(fmt.Sprintf("Total operator '+' of %d and %d is %d", number1, number2, value))

	//Operator minus
	value = number1 - number2
	fmt.Println(fmt.Sprintf("Total operator '-' of %d and %d is %d", number1, number2, value))

	//Operator multiple
	value = number1 * number2
	fmt.Println(fmt.Sprintf("Total operator '*' of %d and %d is %d", number1, number2, value))

	//Operator divider
	value = number1 / number2
	fmt.Println(fmt.Sprintf("Total operator '/' of %d and %d is %d", number1, number2, value))

	//Operator modulus
	value = number1 % number2
	fmt.Println(fmt.Sprintf("Total operator 'modulus' of %d and %d is %d", number1, number2, value))

	/**
		* OPERATOR Perbandingan
	*/
	isTrue := number1 > number2
	fmt.Println("is Number1 more than Number 2, ",isTrue)
	
	/**
		* OPERATOR Logika
	*/
	a := true
	b := false
	
	c := a && b
	fmt.Println("is a == b, ",c)
	
}