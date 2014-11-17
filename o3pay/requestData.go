package o3pay 

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  "bytes"
  "encoding/json"
)
   
type DataPay struct{}

func RequestData(data DataPay,url string)(err error){
     b,err:=json.Marshal(data)
  if(err!=nil){
      fmt.Println("json err:",err)
   }


  body:=bytes.NewBuffer([]byte(b))
  res,err:=http.Post(url,"application/json;charset=utf-8",body)
  if(err!=nil){
    log.Fatal(err)
    return
   }
  result,err:=ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err!=nil{
      log.Fatal(err)
      return
  }
  fmt.Printf("ok----%s",result)
  return
}
