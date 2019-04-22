package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 初始化通道
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //开启goroutine
	}
	for range os.Args[1:] {
		fmt.Println("from ch", <-ch) // 输出ch中的内容
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 将err传递到管道中
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard可以当做垃圾桶 ieru一些不需要的数据
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("whie reading %s:%v\n", url, err)

	}
	secs := time.Since(start).Seconds()
	// fmt.Printf("%.2fs %7d %s\n", secs, nbytes, url) //
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url) // 将结果输出到通道中

}
