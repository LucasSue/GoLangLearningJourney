package main

import "fmt"

type MyBook struct {
	name string
	num  int
}

func (m *MyBook) SetMybook(num int, name string) {
	m.num = num
	m.name = name
}

func (m *MyBook) ShowMyBook() {
	fmt.Printf("My Book is %s,number is %d", m.name, m.num)
}

func ShowMyName() {
	fmt.Println("Lucas")
}

func main() {
	mybook := MyBook{"su", 10}
	mybook.ShowMyBook()
}
