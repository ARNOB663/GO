# Variables and Data Types in Go

## Variables in Go
In Go, variables are explicitly declared and used by the compiler to check the type-correctness of function calls.

### Declaring Variables
Variables can be declared in two ways:
1. Using the `var` keyword:
    ```go
    var name string
    var age int
    var isStudent bool
    ```

2. Using shorthand notation (only inside functions):
    ```go
    name := "John"
    age := 25
    isStudent := true
    ```

### Multiple Variable Declaration
You can declare multiple variables at once:
```go
var name, city string
var age, grade int