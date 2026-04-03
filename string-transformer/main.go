
       // CodeCrafters — Operation Gopher Protocol
       // Module: String Transformer
       // Author: [Joy Adikwu]
       // Squad:  [THE DEPLOYABLES]


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var output string

	fmt.Println("WELCOME TO THE TRANSFORMATION OF STRINGS")

	for {
		fmt.Println("Enter Transformals(cap, lower, upper, title, snake, reverse):")
		scanner.Scan()
		Transformals := strings.TrimSpace(scanner.Text())

		fmt.Println("Enter Strings to Transform:")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		if text == ("quit") {
			fmt.Println("Goodbye! Have a nice day")
			break
		}

		if Transformals == "cap" {
			output = Cap(text)
		} else if Transformals == "upper" {
			output = Upper(text)

		} else if Transformals == "lower" {
			output = Lower(text)
		} else if Transformals == "title" {
			output = Title(text)
		} else if Transformals == "snake" {
			output = Snake(text)
		} else if Transformals == "reverse" {
			output = Reverse(text)
		} else {
			fmt.Println("UNKOWN TRANSFORMATION")
			continue
		}
		fmt.Println("Result:", output)

	}

}

func Upper(text string) string {
	return strings.ToUpper(text)
}

func Lower(text string) string {
	return strings.ToLower(text)
}

func Cap(text string) string {
	return strings.Title(strings.ToLower(text))
}

func Title(text string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true,
		"or": true, "for": true, "nor": true, "on": true, "at": true,
		"to": true, "by": true, "in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}

	words := strings.Fields(strings.ToLower(text))

	for i := 0; i < len(words); i++ {

		if i == 0 || !smallWords[words[i]] {
			words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]
		}
	}

	return strings.Join(words, " ")
}

func Snake(text string) string {

	text = strings.ToLower(text)

	var result []rune

	for _, r := range text {

		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			result = append(result, r)
		} else if r == ' ' {
			result = append(result, '_')
		}
	}

	return string(result)
}

func Reverse(s string) string {
	reverse := func(s string) string {
		r := []rune(strings.TrimSpace(s))
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}

	return strings.Join(words, " ")
}
