package main
import (
  "fmt"
  "os"
  "strconv"
)
func main(){
 /*  p,s,sep :="","",""
   for i,arg := range os.Args[0:]{
       a, b := strconv.Atoi(arg)
        if b==nil{
        t := strconv.Itoa(i)
        p += sep + t
        s += sep + arg
        sep = " "
}
    fmt.Println("Enter an integer index")*/
    
 
   for i,arg := range os.Args[0:]{
      t := strconv.Itoa(i)
      fmt.Println("index"+" "+ t + "    " + "value"+" " + arg )
}  
    /*fmt.Println(p)
    fmt.Println(s)*/
}

