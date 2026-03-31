```markdown

# Condition

## What I learnt
Conditions in go perform different actions base on the condition given. it can either be "true" or "false". it support the operator system.

**It has four condition statements:**
The "if" statement: it execute a block of code, and specify if a condition is true.

The "else" statement: it execute a block of code also but specify if the condition given is false.

The "else if" statement: it checks if the first condition is false.

The "nested if" statement: the are inside if statement. 

## Example

```go
age := 18

if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
