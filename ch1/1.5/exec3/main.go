// 如果输入的url中没有前缀 http:// ,使用strings.HasPrefix
// 获取resp.Status 的状态吗
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		// HasPrefix 返回bool
		if !strings.HasPrefix(url, prefix) {
			url = "http://" + url
		}
		resp, err := http.Get(url) // 结构体resp 得到访问的请求结果 resp.Body中包含可读的服务器响应流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body) // 使用io.Copy 将响应体的内容拷贝到os.Stdout 标准输出
		fmt.Printf("resp status is: %v\n", resp.Status)
		resp.Body.Close() //关闭resp Body流 防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %v err\n", err)
			os.Exit(1)
		}
	}
}

/*
 ./exec1 http://www.baidu.com
*/
