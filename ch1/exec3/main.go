//测试 += 拼接和 strings.Join 的性能
//通过对比在数量是100000的时候 使用拼接的时间是1.62418s,使用strings.Join的时间0.00282s
//是575.95035461倍
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < 100000; i++ {
		sep = strconv.Itoa(i) //将数字转换成字符
		s += sep
		sep = "" //1.62418s
	}

	// var sep []string
	// for i := 0; i < 100000; i++ {
	// 	sep = append(sep, strconv.Itoa(i))
	// 	s := strings.Join(sep[:], "")
	// 	sep = sep[0:0]  //清空切片
	// 	_ = s  //0.00282s
	// }
	fmt.Printf("%.5fs used\n", time.Since(start).Seconds())
}
