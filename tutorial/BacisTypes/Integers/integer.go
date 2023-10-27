package main

import (
  "fmt"
  "strconv"
  "log"
)
func main() {
  //zero value of integer is 0.
  //Convert int to string with strconv.Ito
  var i1 int = -38
  fmt.Printf("i1: %s\n", strconv.Itoa(i1))

  var i2 int = 148
  fmt.Printf("i2: %s\n", strconv.Itoa(int(i2)))

  //Convert int to string with fmt.Sprintf
  s1 := fmt.Sprintf("%d", i1)
  fmt.Printf("i1: %s\n", s1)
  
  s2 := fmt.Sprintf("%d", i2)
  fmt.Printf("i2: %s\n", s2)

  //Convert string to int with strconv.Atoi 
  s := "-48"
  i1, err := strconv.Atoi(s)
  if err != nil {
    log.Fatalf("strconv.Atoi() failed with %s \n", err)
  }
  fmt.Printf("i1: %d\n", i1)

  //Convert string to int with fmt.Sscanf
  s = "348"
  var i int
  _, err2 := fmt.Sscanf(s, "%d", &i)
  if err2 != nil {
    log.Fatal("fmt.Sscanf failde with '%s'\n", err)
  }
  fmt.Printf("i1: %d\n", i)
}
