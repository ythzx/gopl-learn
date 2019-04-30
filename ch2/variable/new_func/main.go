//new 函数
package main

import "fmt"

func main() {
	p := new(int)   // p *int 类型 指向匿名 int变量
	fmt.Println(*p) //初始化为为0
	fmt.Println(p)  //返回的变量的地址
	*p = 2
	fmt.Println(*p)

	a := new(int)
	b := new(int)
	fmt.Println(a == b) //false 每次调用new 函数都返回一个新的地址
}

/*
如下的的结果一样
*/
func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}
