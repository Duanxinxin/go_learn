// @Time    : 2023/12/5 20:39
// @File    : 10-server3.go
// @User    : Recci
// @Software: GoLand

// 更完整的服务器例子

package main

import (
	"fmt"
	"log"
	"net/http"
)

// 处理程序回显请求 URL 的路径部分
func handlerV3(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	if err != nil {
		return
	}
	for k, v := range r.Header {
		_, err = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
	_, err = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if rErr := r.ParseForm(); rErr != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, err = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handlerV3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
