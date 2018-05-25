package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
  file,err:=ioutil.ReadFile("file1.txt")
  if err !=nil {
fmt.Println("Error is ",err)
  }
  var count int
  str:=string(file)
  fmt.Println(strings.ToUpper(str))  
  a:=strings.Split(str, " ")
  count1:=len(a)
  for i:=0;i<len(a);i++ {
      for _,v:=range a[i] {
          switch v {
            case 'a':
            count++
            case 'A':
            count++
            case 'e':
            count++
            case 'E':
            count++
            case 'i':
            count++
            case 'I':
            count++
           case 'o':
            count++
            case 'O':
            count++
            case 'u':
            count++
            case 'U':
            count++
          }
      }
  }
  fmt.Println(count1,count)
}