package main

import (
  "io/ioutil"; 
  )


func main() {
  contents,_ := ioutil.ReadFile("input.txt")
  println(string(contents))
  ioutil.WriteFile("filename", contents, 644)
}
