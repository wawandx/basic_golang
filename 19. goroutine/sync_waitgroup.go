package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	//add //wait //done

	//cara 1
	/* waitGroup.Add(1)
	go printText("Salam", &waitGroup)

	waitGroup.Add(1)
	go printText("Hallo", &waitGroup) */

	//cara 2
	waitGroup.Add(2)
	go printText("Salam", &waitGroup)
	go printText("Hallo", &waitGroup)

	waitGroup.Wait()
}

func printText(text string, waitGroup *sync.WaitGroup) {
	for i:=0; i<5; i++ {
		fmt.Println(text)
	}

	waitGroup.Done()
}