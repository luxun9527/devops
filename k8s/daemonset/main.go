package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := flag.String("port", "12345", "端口")
	flag.Parse()
	log.Printf("Server started on port %s", *port)
	if err := http.ListenAndServe(":"+*port, http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

		// 设置响应头
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 设置响应状态码
		writer.WriteHeader(http.StatusOK)
		// 设置响应体
		writer.Write([]byte(fmt.Sprintf(`{"code":200,"data":{"time":"%v"}}`, time.Now().Format(time.DateTime))))
	})); err != nil {
		log.Panicf("启动服务器失败：%v", err)
	}
}
