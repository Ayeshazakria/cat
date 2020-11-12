
package main
 import(
"fmt"
"math"
 //"os"
)
func main(){
  var i int
  fmt.Print("Enter a number to find its square_root\n")
  fmt.Scanf("%d",&i )
 //fmt.FPrintf(os.Stdout,"%#X\n",i)
 
 //_, err := fmt.Scanf("%d",&i)
  
 //var s float64 =math.Sqrt(i)
 var s float64 =math.Sqrt(float64(i))
  fmt.Println("the square_root of your number is\n",s)
}

