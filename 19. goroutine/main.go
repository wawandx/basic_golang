package main

import (
	"fmt"
	"time"
)

func printText() {
	for i:=0; i<5; i++ {
		fmt.Println("text", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printNumber() {
	for i:=0; i<5; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	start := time.Now()

	go printNumber()
	go printText()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println(time.Since(start))
}