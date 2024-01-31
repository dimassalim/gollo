package main

import (
    "fmt"
    "github.com/dimassalim/go-greet"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}