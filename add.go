package main
import (
  "fmt"
  
"os"
)
func main() {
var sum ,n1,n2 int
fmt.Println("Enter first integer to get its sum")
fmt.Scanf("%d",&n1)
fmt.Println("Enter second integer to get its sum")
fmt.Scanf("%d",&n2)
 fmt.Fprintf(os.Stdout, "%#X\n",  "%#X\n", n1,n2)

sum=n1 +n2
fmt.Println(sum)}
