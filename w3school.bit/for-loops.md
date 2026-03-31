```markdown

# For loop

## What i learnt
for loop loops through or iterate a block of code many times based on the input given.

statement in loop:
the continue statement: The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

the break statement: The break statement is used to break/terminate the loop execution.

nested loop: the "inner loop" will be executed one time for each iteration of the "outer loop":

Range keyword: The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

## Example
```go

 package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}