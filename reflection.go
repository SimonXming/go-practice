package main

import (
	"fmt"
)

type MyStruct struct {
	name string
}
type MyStruct2 struct {
	name string
}

func (this *MyStruct) GetName(str string) string {
	this.name = str
	return this.name
}

type ReflectionInterface interface {
	GetName(string) string
}

func main() {
	var accomplishObj0 MyStruct = MyStruct{"Tom"}
	// it's will raise a panic
	// panic: interface conversion: main.MyStruct is not main.ReflectionInterface: missing method GetName
	fmt.Println(interface{}(accomplishObj0).(ReflectionInterface))

	var emptyInterface interface{}

	var value interface{}
	switch str := value.(type) {
	case string:
		fmt.Println(str, "string")
	case ReflectionInterface:
		fmt.Println(str, "ReflectionInterface")
	default:
		fmt.Println("Something else.")
	}

	accomplishObj1 := interface{}(new(MyStruct))
	fmt.Println(accomplishObj1.(ReflectionInterface))

	// it's will raise a panic
	// panic: interface conversion: *main.MyStruct2 is not main.ReflectionInterface: missing method GetName
	accomplishObj2 := interface{}(new(MyStruct2))
	fmt.Println(accomplishObj2.(ReflectionInterface))

	// it's will NOT raise a panic
	// <nil> false
	var i1, ok = interface{}(emptyInterface).(ReflectionInterface)
	fmt.Println(i1, ok)

	// it's will raise a panic
	// panic: interface conversion: interface is nil, not main.ReflectionInterface
	fmt.Println(emptyInterface.(ReflectionInterface))
}
