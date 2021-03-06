package main

import (
	"fmt"
	"time"
)

func main() {
	goroutine_closures_1()
	goroutine_closures_2()
}

func goroutine_closures_1() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}
	fmt.Println("out of loop")

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
	fmt.Println("out of func goroutine_closures_1")

}

// goroutine 的执行时机是不确定的
func goroutine_closures_2() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
		// time.Sleep(time.Millisecond)
	}
	fmt.Println("out of loop")
	time.Sleep(time.Millisecond)

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
	fmt.Println("out of func goroutine_closures_2")

}
