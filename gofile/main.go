package main

import "fmt"

func PrintSomeThing() {
	fmt.Println("This is my first go project")
}

func main() {
	mybook := MyBook{"su", 10}
	mybook.ShowMyBook()
}
