package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // 结构体resp 得到访问的请求结果 resp.Body中包含可读的服务器响应流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) //ioutil.ReadAll 从response 中读取全部内容
		resp.Body.Close()                   //关闭resp Body流 防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %v err\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

/*
 ./fetch http://www.baidu.com
*/
