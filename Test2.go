package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var certPath = "ca/server.crt"
var keyPath = "ca/server.key"

func main() {
	http.HandleFunc("/gohttps", httpsHandler)
	http.ListenAndServeTLS(":8080", certPath, keyPath, nil)
}

func httpsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	contentType := r.Header.Get("Content-Type")
	fmt.Println("contentType", contentType)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))
	// 回复
	//w.Write([]byte("www.5lmh.com"))
	//w.Write([]byte("1"))
	m := []string{"AA", "BB"}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(m)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}
