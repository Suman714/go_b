package main

import (
  "fmt"
  "unsafe"
)

func main() {
  var b bool = true
  fmt.Printf("b is: %v \n", b)
  b = false
  fmt.Printf("b is: %v \n", b)
  var b2 bool
  fmt.Printf("Zero value of bool is: %v \n", b2)
  fmt.Printf("size of bool is: %d\n", unsafe.Sizeof(b))
}
