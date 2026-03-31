package main

import (
	"fmt"
)

func main() {
	var firstnum int
	var lastnum int
	var operator string

	start:
    fmt.Println("welcome! to my calculator")
	fmt.Println("Enter firstnum:" )
	_, err1 := fmt.Scan(&firstnum)
	fmt.Println("Enter lastnum:" )
	_, err2 := fmt.Scan(&lastnum)

	if err1 != nil || err2 != nil {
		fmt.Println("Input a valid number")
		goto start
	}
	fmt.Println("Enter operator:" )
	fmt.Println("add")
	fmt.Println("sub")
	fmt.Println("mul")
	fmt.Println("div")
	fmt.Println("quit")
	fmt.Println("help")
	fmt.Scan(&operator)

	if operator == "add" {
			fmt.Println(firstnum+lastnum)
    }
	if operator == "sub" {
		fmt.Println(firstnum-lastnum)
	}
	if operator == "mul" {
		fmt.Println(firstnum*lastnum)
	}
	if operator == "div" {
		if lastnum == 0 {
			fmt.Println("invalid can not divide by zero")
		}else {
			fmt.Println(firstnum/lastnum)
		}
	}
	if operator == "quit" {
		fmt.Println("thanks for using my calculator")
		return 
	}
	if operator == "help" {
		fmt.Println("Help instruction:" )
		fmt.Println("Enter firsnum first")
		fmt.Println("Enter lastnum next")
		fmt.Println("Enter operator")
		fmt.Println("Your answer wiil be give after you choose your operator")
		fmt.Println("if you enter quit a goodby message wiil be displayed")
	}
	main()
}