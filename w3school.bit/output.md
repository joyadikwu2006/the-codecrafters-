```markdown
# Output

## What I learnt 
**Go output functions**
they are three output text in go;
*Print(): this give the default format of the argument.
*Println(): it used to print the output in the console in a newlines.
*Printf() : this is for printing the formatting verb of the argument.

**Go formatting verbs**
formatting verb can be used with printf()
# Examples of formatting verb
* %v  Print the value in the default format
* %#v Print value of the argument in go syntax format
* %T Print the argument type
* %s print the value as a string
e.t.c.

# Example
```go
package main

import "fmt"

func main() {
    s := "Diamond"
    fmt.Print(s): this gives you the output "Diamond" without a newline.
    fmt.Println(s): this will give you "Diamond" but with a newline after printing the argument.
    fmt.Printf("%c", s): this will give something "!c(string=diamond)" give you the type it is.
}