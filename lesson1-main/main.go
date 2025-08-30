// hello.go
// package declaration
package main

// import package
import "fmt"

// function
func add(a, b int) int {
  return a+b
}
// global variable
var g int = 100

func main() {
  a, b := 1, 2
  res := add(a, b)
  fmt.Println("a=", a, "b=", b, "a+b=", res)
  fmt.Println("g=", g)
  fmt.Println("hello world!")
  fmt.Println(g + res)
}