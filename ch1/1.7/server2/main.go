//对请求次数进行计算
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//这里会有两次请求 ？？ 是浏览器中自动添加 /  如果使用/test 访问/count 就不会请求/test
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("come /")
	mu.Lock() // 每个请求都会是一个goroutine 加锁保证同一时间只有一个goroutine 处理
	count++
	fmt.Fprintf(w, "Count:%d\n", count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//访问/count 会调用handler
func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("come /count")
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
