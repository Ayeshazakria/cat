package main

import "fmt"
import "flag"

func main(){

   var  result int
    
   v:=1

   fmt.Println(v)
   
   var  p *int

 // var *p int  // syntax error

   fmt.Println(p) // <nil>
  
   p = &v

   fmt.Println(p) // address of v
   
   result = showAll(p)

   fmt.Println("Now the updated value of v is :" , result)
   fmt.Println(o)
   fmt.Println(sep)   
 }
  
 func showAll(z *int) int  {



   fmt.Println(z)    //address of v stored in the copy of pointer p (which is passed as argument)

   fmt.Println(*z)   // print 1   (*z prints the value to which pointer z is pointing,that is p,which further holds v)

//   z=5 // syntax error

    *z = 5
    
   fmt.Println(*z==5)
    
    *z++
    
   fmt.Println(*z)
   
   return *z    
    
    
 }
 var o =flag.Bool("o",false,"omit new line")
 var sep =flag.String("s", "/", "seperator")


