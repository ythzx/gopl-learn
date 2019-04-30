package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //三个参数：命令行标志位名字 默认参数 描述信息 -n 忽略行尾换行符
var sep = flag.String("s", " ", "separator")           //字符串类型标志参数 -s 指定分隔符 默认是空格

//n sep 必须使用指针取值 分别对应命令行参数 n s对应的指针
func main() {
	flag.Parse()                               // 首先执行解析 用于更新每个标志参数对应的变量的值
	fmt.Print(strings.Join(flag.Args(), *sep)) //拼接字符串 这里用的是sep的值(string) flag.Args()返回值是对应类型的slice
	//加上-n 参数，不执行下面语句
	fmt.Printf("the default *n is :%v\n", *n)
	// 添加-n参数后 *n 的值为true
	if !*n {
		fmt.Printf("the *n is :%v\n", *n)
		fmt.Println()
	}

}
