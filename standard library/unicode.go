package main

import (
    "fmt"
)

func main(){
    var s string = "中文测试"
    fmt.Println(s[1])
    var t = []rune(s)
    fmt.Println(t)
    fmt.Printf("%c %c", t[0], t[1])
}
