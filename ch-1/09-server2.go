// @Time    : 2023/12/5 20:32
// @File    : 09-server2.go
// @User    : Recci
// @Software: GoLand

// 计数器服务器

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// 处理程序回显请求 URL 的路径部分
func handlerV2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, fprintf := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if fprintf != nil {
		return
	}
}

// counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Count %d\n", count)
	if err != nil {
		return
	}
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handlerV2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
