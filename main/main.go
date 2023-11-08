package main

import (
	"GoLangLearningJourney/test1"
	"fmt"
)

func PrintSomeThing() {
	fmt.Println("This is my first go project")
}

func main() {
	mybook := MyBook{"su", 10}
	mybook.ShowMyBook()
	test1.ShowMyName()
	test1.ShowYourName()
}
