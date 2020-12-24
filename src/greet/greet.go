package greet

import "fmt"

var Greeting = "Hello"

func Hello(name string) string{
  return fmt.Sprintf("%s %s", Greeting, name)
}



