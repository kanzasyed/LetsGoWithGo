package main

import (
    "fmt"
    "LetsGoWithGo/greetings"
)
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
  fmt.Println(quote.Go())
  // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
