package main

import (
	"fmt"
	"myproject/Task1"
	"myproject/Task2"
	"myproject/Task3"
	"myproject/Task4"

	"github.com/fatih/color"
)

func main() {
	defer color.Red("Exit")
	defer color.Green("This app will be updated weekly or more frequently. Check your version at https://github.com/Erdxd/taskongolang.")

	var menu = map[string]func(){
		"1": Task1.Task1,
		"2": Task2.Task2,
		"3": Task3.Task3,
		"4": Task4.Task4,
	}
menu:
	for {
		Useranswer := Enter([]string{
			"1-Task 1",
			"2-Task 2",
			"3-Task 3",
			"4-Task 4",
			"0 - close",
			"Выберите вариант",
		})
		FuncMenu := menu[Useranswer]
		if Useranswer == "0" {

			break menu
		} else if FuncMenu == nil {
			color.Yellow("Error, please repeat attempt")
			continue
		}

		FuncMenu()

	}
}
func Enter[T any](prompt []T) string {

	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scanln(&res)
	return res

}
