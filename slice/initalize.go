package main

import "fmt"

func InitializeArray() [4]int {
	number := [4]int{0, 1, 3, 4}
	for i := 0; i < len(number); i++ {
		number[i] = i * 10
	}
	return number
}

func main() {
	fmt.Println(InitializeArray())
}
