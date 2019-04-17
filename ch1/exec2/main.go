//练习2 打印每个参数的索引和值 每个为一行
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Print(index, arg)
		fmt.Println()
	}
}
