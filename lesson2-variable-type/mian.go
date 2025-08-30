package main

import "fmt"

func main() {
    str := "abc"
    str = "def" // ok
    /* 下面的就不行，编译报错：cannot assign to str[0] (strings are immutable)
    str[0] = "d"金矛勇者蜡笔小新 金矛勇者
    */
    fmt.Println(str)
}