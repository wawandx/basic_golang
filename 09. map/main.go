package main

import (
	"fmt"
)

func main() {
	var mahasiswa = map[string]string{}
	//var mahasiswa = map[string]string{"nama": "Rachmad Kurniawan", "alamat": "Kembangan"}

	mahasiswa["nama"] = "Rachmad Kurniawan"
	mahasiswa["alamat"] = "Kembangan"
	fmt.Println(mahasiswa)

	for i, v := range mahasiswa {
		fmt.Println(i, "=", v)
	}

	delete(mahasiswa, "alamat")
	for i, v := range mahasiswa {
		fmt.Println(i, "=", v)
	}

	var dataMahasiswa = []map[string]string{
		map[string]string{"nama": "Rachmad Kurniawan", "alamat": "Kembangan"},
		map[string]string{"nama": "Riski Anggita", "alamat": "Palu"},
	}
	
	for i, v := range dataMahasiswa {
		fmt.Println(i, "= nama :", v["nama"], ", alamat:", v["alamat"])
	}
}