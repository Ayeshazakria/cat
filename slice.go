package main

import "fmt"

func main(){
 arr:=[...]string{"mon","tue","wed","thur"}
 fmt.Println("Array",arr)
 
 sli:=arr[0:4]
 fmt.Println("slice",sli)

 if len(arr)==len(sli){

    fmt.Println("having same lenght")
 }
 
 }
