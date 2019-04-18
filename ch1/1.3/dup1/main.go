//查找重复的行
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //map 是集合 k-v 使用make初始化
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf("The input is: %s\n", input.Text())
		counts[input.Text()]++ // counts[line] = counts[line] + 1 counts[line] 初始值为0
 		fmt.Printf("The counts is: %v\n", counts)

		fmt.Println("数量----内容")
		//map 中元素的顺序是无序的
		for key, value := range counts {
			if value > 1 {
				fmt.Printf("%d\t%s\n", value, key)
			}
		}
	}

}

/*
The input is: aa
The counts is: map[cc:1 aa:2 bb:1]
数量----内容
2       aa
bb
The input is: bb
The counts is: map[aa:2 bb:2 cc:1]
数量----内容
2       aa
2       bb
aa
The input is: aa
The counts is: map[bb:2 cc:1 aa:3]
数量----内容
3       aa
2       bb
*/
