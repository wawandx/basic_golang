package main

import (
	"fmt"
)

func main() {
	var mp interface{}

	data := map[string]string{
		"Name":    "Rachmad Kurniawan",
		"Address": "Kembangan, Jakarta Barat",
		"Website": "Agesa.co.id",
	}

	mp = map[string]interface{}{
		"data": data,
	}

	fmt.Println(mp)
}