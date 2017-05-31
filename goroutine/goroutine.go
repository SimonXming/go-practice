package main

import (
	"fmt"
)

/*
channel <- value // 阻塞发送
<- channel // 接收并将其丢弃
x := <- channel // 接收并将其保存
x, ok := <- channel // 接收并将其保存，同时检查通道是否已关闭或者是否为空
*/

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}

func main() {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A->%d, B->%d)", a, <-counterB)
	}
	fmt.Println()
}
