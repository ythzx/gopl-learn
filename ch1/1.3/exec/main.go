//read the file and count the duplicate lines
//打印重复行的文件的名字
package main

import (
	"bufio"
	"fmt"
	"os"
)

var filename []string

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("file is none")
		// countLines(os.Stdin, counts) //传入一个输入对象
	} else {
		for file, arg := range files {
			fmt.Printf("the file is: %v\n", file)
			f, err := os.Open(arg) // 函数返回的第一个是 *os.File
			// f, err := os.Open("./1.txt") // 函数返回的第一个是 *os.File
			fmt.Printf("file obj f is: %v\n", f)         //文件对象
			fmt.Printf("arg is: %v\n", arg)              //读取的文件的名字
			fmt.Printf("os.arg[0] is: %v\n", os.Args[0]) // 程序自己的名字
			filename = append(filename, arg)             // 切片再添加后需要重新赋值
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err) //Fprintf 打印文件的错误
				continue
			}
			countLines(f, counts, filename)
			f.Close()
		}

	}
	fmt.Printf("the counts is: %v\n", counts)
	fmt.Print("counts\tcontent\n")
	for key, value := range counts {
		if value > 1 {
			fmt.Printf("%d\t%s\t%s\n", value, key, filename[:])
		}
	}

}

/*
hzx:dup2 mac$ ./dup2 1.txt 2.txt
the file is: 0
the file is: 1
the counts is: map[222:4 text1:2 111:6]
counts  content
2       text1
6       111
4       222
hzx:dup2 mac$
*/

func countLines(f *os.File, counts map[string]int, arg []string) {
	input := bufio.NewScanner(f) //
	for input.Scan() {
		counts[input.Text()]++
	}
}

/*
// 将代码修改成打印重复的文件名字
os.arg[0] is: ./dup2
the counts is: map[text1:2 111:6 222:4]
counts  content
2       text1   [1.txt 2.txt]
6       111     [1.txt 2.txt]
4       222     [1.txt 2.txt]
*/
