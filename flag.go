package main

import (
  "flag"
  "fmt"
  "strings"
)
 var n = flag.Bool("a", true,"omit newlines")
 var sep=flag.String("s", "_ ", "seperator")
 
 func main() {
   flag.Parse()
   fmt.Print(strings.Join(flag.Args(),*sep))
   if !*n {
     fmt.Println()
}
}
