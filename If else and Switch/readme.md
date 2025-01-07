# README.md

## Conditional Statements in Go: `else if` and `switch`

### `else if` in Go

The `else if` statement in Go is used to check multiple conditions sequentially. If one condition evaluates to `true`, its associated block of code will execute, and the remaining conditions are ignored. 

#### Syntax:
```go
if condition1 {
    // Code block for condition1
} else if condition2 {
    // Code block for condition2
} else {
    // Code block if none of the conditions are true
}
```

#### Example:
```go
package main

import "fmt"

func main() {
    x := 20

    if x < 10 {
        fmt.Println("x is less than 10")
    } else if x == 20 {
        fmt.Println("x is exactly 20")
    } else {
        fmt.Println("x is greater than 10 but not 20")
    }
}
```

### `switch` in Go

The `switch` statement in Go is used to simplify multiple conditional checks. It evaluates an expression and executes the first matching `case` block. Unlike other languages, Go's `switch` does not require explicit `break` statements, as they are implicit.

#### Syntax:
```go
switch expression {
case value1:
    // Code block for value1
case value2:
    // Code block for value2
default:
    // Code block if no cases match
}
```

#### Example:
```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("Start of the work week!")
    case "Friday":
        fmt.Println("Almost the weekend!")
    default:
        fmt.Println("It's just another day.")
    }
}
```

### Key Differences Between `else if` and `switch`:

| Feature               | `else if`                          | `switch`                           |
|-----------------------|------------------------------------|------------------------------------|
| **Usage**            | For sequential condition checks   | For evaluating a single expression|
| **Code readability** | Can get verbose with many checks  | Cleaner and more organized        |
| **Default behavior** | Explicit `else` block required    | Implicit `default` case optional  |

---

### How to Run the Code
1. Install Go from [golang.org](https://golang.org/).
2. Save the example code in a `.go` file (e.g., `main.go`).
3. Open a terminal and navigate to the file's directory.
4. Run the code using the command:
   ```bash
   go run main.go
   
