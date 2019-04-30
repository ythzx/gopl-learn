//函数中返回局部变量的地址
//每次调用f() 都会返回不同的地址
package main

import "fmt"

func main() {
	var p = f()
	fmt.Printf("the p is: %v\n", p)
	// fmt.Println("compare the two f()", f() == f())
	fmt.Println("first f()", f())
	fmt.Println("second f()", f())

}

func f() *int {
	v := 1 // 局部变量v 在
	// fmt.Printf("the func f() return v is: %v\n", v)
	fmt.Printf("the func f() return addr is: %v\n", &v)
	return &v
}
