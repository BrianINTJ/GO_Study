package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<type:Book , value:book{}地址>
	b := &Book{}

	// r: pair<type: , value:>
	var r Reader
	// r: pair<type:Book , value:book{}地址>
	r = b

	r.ReadBook()
	
	// w: pair<type: , value:>
	var w Writer
	// w: pair<type:Book , value:book{}地址>
	w = r.(Writer)

	w.WriteBook()
}
