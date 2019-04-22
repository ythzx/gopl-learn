// 使用io.Copy 将响应体的内容拷贝到os.Stdout (标准输出) 避免申请一个缓冲区来存储
package main

import (
	"fmt"
	"io"
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
		_, err = io.Copy(os.Stdout, resp.Body) // 使用io.Copy 将响应体的内容拷贝到os.Stdout 标准输出
		resp.Body.Close()                      //关闭resp Body流 防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %v err\n", err)
			os.Exit(1)
		}
	}
}

/*
 ./exec1 http://www.baidu.com
*/
