package main

import (
	"fmt"
	"myproject/Task1"
	"myproject/Task2"
)

func main() {
	for {
		fmt.Println("Enter the number of the problem you want to get an answer to.(1-2)")
		var User string
		fmt.Scan(&User)
		switch User {
		case "1":
			Task1.Task1()
		case "2":
			Task2.Task2()
		}
	}
}
