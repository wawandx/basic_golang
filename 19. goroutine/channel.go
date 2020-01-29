package main

import (
	"fmt"
	"time"
)

var itemsChannel = make(chan string)

func main() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "harta", "kerang"}
	// casenya : 
	// Kita mau menyelam mencari harta karun,
	// terselip diantara banyak barang
	// 1. penyelam, 2. bersihkan, 3. Disimpan
	
	go menyelam(items)
	go membersihkan()

	time.Sleep(500 * time.Millisecond)
}

func menyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil mendapatkan " + item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil membersihkan " + rawItem)
	}
}
