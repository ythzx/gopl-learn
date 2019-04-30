package main

import "fmt"

func main() {
	var x int
	p := &x // point p 执行x的int类型的指针 指针保存了变量x的内存地址
	fmt.Printf("the point p:%v\n", p)

	fmt.Printf("the point p value is:%v\n", *p) // *p读取指针所指向的变量的值
	*p = 2
	fmt.Printf("the point p update value is:%v\n", *p)
	fmt.Printf("the x value is:%v\n", x)

	var a, b int
	fmt.Println(&a == &a, &a == &b, &a != nil) // 指针测试 a b 在声明的时候已经赋 类型零值
}
