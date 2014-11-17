package main

import "fmt"


func main(){
 var student map[string]string
 student=make(map[string]string,5)

 student["stu1"]="Peter"
 student["stu2"]="Anne"
 student["stu3"]=""
 
 fmt.Printf("%v\n",student)
 
 for _,value:=range student{
   if value!=""{
   fmt.Println(value)
 }
}
}
