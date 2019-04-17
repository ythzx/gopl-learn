package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// var s1, sep1 string
	// for _, arg := range os.Args[1:] {
	// 	s1 += sep1 + arg
	// 	sep1 = " "
	// }
	// fmt.Println(s1)

	// fmt.Println(strings.Join(os.Args[1:], " ")) //使用strings Join可以拼接数据量大的字符串
	fmt.Println(os.Args[1:])
}
