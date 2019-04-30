//将指针作为参数调用函数，可以在函数中通过该指针更新该变量的值
package main

import "fmt"

func main() {
	v := 1
	// incr(&v)
	fmt.Printf("the &v addr  is: %v\n", &v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	fmt.Printf("the point parm is: %v\n", p)
	fmt.Printf("the point parm value is: %v\n", *p)
	*p++ // 仅仅是更新指针p指向的值 不是指针p本身
	return *p
}
