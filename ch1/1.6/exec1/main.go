//连续对同一个网站进行请求并比较返回的数据是否有差异
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
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard可以当做垃圾桶 存储一些不需要的数据
	// nbytes, err := io.Copy(os.Stdout, resp.Body)

	/*
		将响应的数据返回并写入到文件中
	*/
	// newFile, err := os.Create("test.html")
	// if err != nil {
	// 	fmt.Println("creat file err")
	// }
	// _, err = io.Copy(newFile, resp.Body)
	// if err != nil {
	// 	fmt.Println("write file err")
	// }

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("whie reading %s:%v\n", url, err)

	}
	secs := time.Since(start).Seconds()
	// fmt.Printf("%.2fs %7d %s\n", secs, nbytes, url) //
	ch <- fmt.Sprintf("%.5fs %7d %s", secs, nbytes, url) // 将结果输出到通道中

}

/*
chanel 是FIFO 先进先出
但是下面的执行是开始耗费的时间少 之后耗费的时间多 ？？？
hzx:exec1 mac$ ./exec1 http://www.baidu.com http://www.baidu.com
from ch 0.0789379770s   153418  http://www.baidu.com
from ch 0.0961079330s   153639  http://www.baidu.com
0.10s elapsed
*/

/*
多次请求会缓存 提高下一次请求的速度
 hzx:exec1 mac$ ./exec1 http://www.v2ex.com http://www.v2ex.com  http://www.v2ex.com http://www.v2ex.com  http://www.v2ex.com  http://www.v2ex.com
from ch 0.66017s   90514 http://www.v2ex.com
from ch 0.73943s   90514 http://www.v2ex.com
from ch 1.17764s   90514 http://www.v2ex.com
from ch 1.19300s   90514 http://www.v2ex.com
from ch 1.19404s   90514 http://www.v2ex.com
from ch 1.20874s   90514 http://www.v2ex.com
*/
