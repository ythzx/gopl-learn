//read the file and count the duplicate lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

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
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err) //Fprintf 打印文件的错误
				continue
			}
			countLines(f, counts)
			f.Close()
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

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f) //
	for input.Scan() {
		counts[input.Text()]++
	}
}
