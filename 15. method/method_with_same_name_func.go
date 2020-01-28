package main

import (
	"fmt"
)

type (
	Category struct {
		Name string
	}

	Post struct {
		Title string
	}
)

func (c Category) lihatData() {
	fmt.Println(c)
}
func (p Post) lihatData() {
	fmt.Println(p)
}

func main() {
	fmt.Printf("From Category\n")
	cats := Category{
		Name: "Berita",
	}

	cats.lihatData()

	fmt.Printf("From Post\n")

	p := Post{
		Title: "Belajar Golang",
	}
	p.lihatData()
}