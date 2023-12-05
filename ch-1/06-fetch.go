// @Time    : 2023/12/4 20:20
// @File    : 06-fetch.go
// @User    : Recci
// @Software: GoLand

// fetch 输出从 URL 获取内容

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // 产生GET请求
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body) // resp 的 Body 域包含 服务器端响应的可读数据流
		closeErr := resp.Body.Close()   // 关闭数据流避免资源泄露
		if closeErr != nil {
			fmt.Printf("请求体关闭异常, %v\n", closeErr)
			return
		}

		if err != nil {
			_, fErr := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if fErr != nil {
				return
			}
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
