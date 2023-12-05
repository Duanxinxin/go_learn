// @Time    : 2023/12/5 20:25
// @File    : 08-server1.go
// @User    : Recci
// @Software: GoLand

// 迷你服务器

package main

import (
	"fmt"
	"log"
	"net/http"
)

// 处理程序回显请求 URL r 的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler) // 请求调用处理程序
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
