//对请求次数进行计算
//打印HTTP请求头和form数据
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
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//这里会有两次请求 ？？
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method:%s URL:%s Proto%s \n", r.Method, r.URL, r.Proto)
	fmt.Fprint(w, "############# Header #############\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=[%q]\n", k, v)
	}
	fmt.Fprintf(w, "Host:%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr:%q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, "############# Form #############\n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
	fmt.Println("come to fist /")
	mu.Lock() // 每个请求都会是一个goroutine 加锁保证同一时间只有一个goroutine 处理
	count++
	fmt.Fprintf(w, "Count:%d\n", count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// http://localhost:8000/test?query=qqq&s=ssss
// Method:GET URL:/test?query=qqq&s=ssss ProtoHTTP/1.1 
// ############# Header #############
// Header["Connection"]=[["keep-alive"]]
// Header["Upgrade-Insecure-Requests"]=[["1"]]
// Header["User-Agent"]=[["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"]]
// Header["Accept"]=[["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]]
// Header["Accept-Encoding"]=[["gzip, deflate, br"]]
// Header["Accept-Language"]=[["zh-CN,zh;q=0.9"]]
// Header["Cookie"]=[["aria2filters=%22%7B%5C%22s%5C%22%3Atrue%2C%5C%22a%5C%22%3Atrue%2C%5C%22w%5C%22%3Atrue%2C%5C%22c%5C%22%3Atrue%2C%5C%22e%5C%22%3Atrue%2C%5C%22p%5C%22%3Atrue%2C%5C%22r%5C%22%3Atrue%7D%22; aria2conf=%7B%22host%22%3A%22localhost%22%2C%22path%22%3A%22/jsonrpc%22%2C%22port%22%3A%226801%22%2C%22encrypt%22%3Afalse%2C%22auth%22%3A%7B%22token%22%3A%22%22%7D%2C%22directURL%22%3A%22%22%7D; session_id=de1b73e26518703836b80e243993175111045a08; Pycharm-c1907b75=9e4b3aaf-4997-44ee-b387-a16884170109"]]
// Host:"localhost:8000"
// RemoteAddr:"127.0.0.1:53353"
// ############# Form #############
// Form["query"]=["qqq"]
// Form["s"]=["ssss"]
// Count:5
// URL.Path = "/test"
