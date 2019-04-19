//read the file and count the duplicate lines
//read all file content to 内存
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	//ioutil 读取文件
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) // 返回字节切片 必须转换为string 才能spilt
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//strings.Split
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	fmt.Printf("the counts is: %v\n", counts)
	fmt.Print("counts\tcontent\n")
	for key, value := range counts {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f) //
	for input.Scan() {
		counts[input.Text()]++
	}
}

/*
hzx:dup3 mac$ ./dup3 1.txt 2.txt
the counts is: map[111:6 222:4 text1:2]
counts  content
2       text1
6       111
4       222
*/
