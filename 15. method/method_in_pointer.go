package main

import (
	"fmt"
)

type (
	Category struct {
		Name string
	}
)

func (c Category) lihatKategori() {
	c.Name = "Berita"
}
func (c *Category) lihatKategoriLagi() {
	c.Name = "Teknologi"
}

func main() {
	cats := Category{
		Name: "Info",
	}

	cats.lihatKategori()
	fmt.Println(cats)

	(&cats).lihatKategoriLagi()
	fmt.Println(cats)
}