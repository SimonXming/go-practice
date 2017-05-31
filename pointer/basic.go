/*
调用 method 的无论是 value 还是 (value's pointer) 都没关系,
method 的定义决定了这个 method 期望得到 value 还是 (value's pointer)
*/

package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) ChangeName(name string) error {
	u.Name = name
	return nil
}

func (u User) ChangeName2(name string) error {
	u.Name = name
	return nil
}

func main() {

	// Value of type User can be used to call the method
	// with a value receiver.
	bill := User{"Bill", "bill@email.com"}
	bill.ChangeName("abc")
	fmt.Println(bill)
	(&bill).ChangeName("bcd")
	fmt.Println(bill)
	bill.ChangeName2("cde")
	fmt.Println(bill)
	(&bill).ChangeName2("def")
	fmt.Println(bill)

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	jill := &User{"Jill", "jill@email.com"}
	jill.ChangeName("abc")
	fmt.Println(jill)
	(*jill).ChangeName("bcd")
	fmt.Println(jill)
	jill.ChangeName2("cde")
	fmt.Println(jill)
	(*jill).ChangeName2("def")
	fmt.Println(jill)
}
