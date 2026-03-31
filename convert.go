package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number string
	var base string

	for {

		fmt.Println("WELCOME TO THE BASE CONVERTER")
		fmt.Print("Enter number: ")

		fmt.Scan(&number)

		if number == "quit" {
			fmt.Println("Goodbye")
			break
		}

		fmt.Println("Enter base: ")
		fmt.Scan(&base)

		switch base {

		case "hex":

			num, err := strconv.ParseInt(number, 16, 64)

			if err != nil {
				fmt.Println("Invalid hexadecimal number")
				continue
			}

			fmt.Println("Decimal:", num)

		case "bin":

			num, err := strconv.ParseInt(number, 2, 64)

			if err != nil {
				fmt.Println("Invalid binary number")
				continue
			}

			fmt.Println("Decimal:", num)

		case "dec":

			num, err := strconv.ParseInt(number, 10, 64)

			if err != nil {
				fmt.Println("Invalid decimal number")
				continue
			}

			fmt.Println("Binary:", strconv.FormatInt(num, 2))
			fmt.Println("Hex:", strconv.FormatInt(num, 16))

		default:

			fmt.Println("Base must be: hex, bin, or dec")

		}
	}
}
