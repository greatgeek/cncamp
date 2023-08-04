package main

import (
	"log"
	"net/http"
	"os"
)

func httpServer() {
	http.HandleFunc("/healthz", healthzHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 复制 request header 到 response header
	for name, headers := range r.Header {
		for _, h := range headers {
			w.Header().Add(name, h)
		}
	}

	// 从环境变量中读取 "VERSION" 配置，并写入 response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))

	// 将客户端的 ip 和 Http 返回码写入日志
	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Proto)

	// 写入状态码 200
	w.WriteHeader(http.StatusOK)
}

func main() {
	httpServer()
}
