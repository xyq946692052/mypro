package main

import (
  "fmt"
  "reflect"
)

type Student struct{
  Name  string
  Age   int
}

func Struct2Map(obj interface{}) map[string]interface{}{
  t:=reflect.TypeOf(obj)
  v:=reflect.ValueOf(obj)
  
  var data=make(map[string]interface{})

  for i:=0;i<t.NumField();i++{

     data[t.Field(i).Name]=v.Field(i).Interface()
  }
  return data
}

func main(){
  stu:=Student{"Peter",25}
  data:=Struct2Map(stu)
  fmt.Println(data)
}
