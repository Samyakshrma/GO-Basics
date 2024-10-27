// 1. Variables and Data Types
package main

import "fmt"

func variablesAndTypes() {
    // Variable declaration methods
    var name string = "John"     // Explicit type
    age := 25                    // Type inference
    var isActive = true         // Type inference with var
    
    // Basic data types
    var (
        integer int = 42                // Integer
        float float64 = 3.14           // Floating point
        character rune = 'A'           // Unicode character
        text string = "Hello, Go!"     // String
        isValid bool = true            // Boolean
    )

    // Constants
    const Pi = 3.14159
    
    // Arrays and Slices
    numbers := [5]int{1, 2, 3, 4, 5}    // Fixed-size array
    slice := []int{1, 2, 3}             // Slice (dynamic size)
    
    // Maps
    person := map[string]string{
        "name": "John",
        "city": "New York",
    }
    
    // Structs
    type User struct {
        Name string
        Age  int
    }
    user := User{Name: "Alice", Age: 30}
}

// 2. Control Structures
func controlStructures() {
    // If-else
    age := 18
    if age >= 18 {
        fmt.Println("Adult")
    } else if age >= 13 {
        fmt.Println("Teenager")
    } else {
        fmt.Println("Child")
    }

    // Switch
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("Weekend soon")
    default:
        fmt.Println("Regular day")
    }

    // For loops
    // Standard for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Range-based for loop
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // While-like for loop
    count := 0
    for count < 5 {
        count++
    }
}

// 3. Functions and Methods
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Method (function with receiver)
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 4. Error Handling
func errorHandlingExample() {
    // Basic error handling
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)

    // Custom error type
    type CustomError struct {
        Code    int
        Message string
    }

    func (e *CustomError) Error() string {
        return fmt.Sprintf("error %d: %s", e.Code, e.Message)
    }
}

// 5. Packages and Imports
// File: mypackage/helper.go
package mypackage

// Exported function (starts with uppercase)
func HelloWorld() string {
    return "Hello, World!"
}

// Unexported function (starts with lowercase)
func internalHelper() string {
    return "Internal helper"
}

// Main program
package main

import (
    "fmt"
    "mypackage"
)

func main() {
    // Using imported package
    message := mypackage.HelloWorld()
    fmt.Println(message)
}

// 6. Defer, Panic, and Recover
func deferExample() {
    // Defer (executes when function returns)
    defer fmt.Println("This prints last")
    fmt.Println("This prints first")

    // Panic and Recover
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong!")
}

// 7. Interfaces
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Usage of interface
func printArea(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
}
