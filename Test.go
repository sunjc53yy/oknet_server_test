//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"oknet-server/pbout"
	"os"

	"google.golang.org/protobuf/proto"
)

var certPath = "ca/server.crt"
var keyPath = "ca/server.key"

func main() {
	http.HandleFunc("/go", httpHandler)
	http.HandleFunc("/goget", getHttpHandler)
	http.HandleFunc("/gofile", fileHandler)
	http.HandleFunc("/gohead", headHttpHandler)
	http.HandleFunc("/goput", putHttpHandler)
	http.HandleFunc("/godelete", deleteHttpHandler)
	http.HandleFunc("/gopb", pbHttpHandler)
	http.HandleFunc("/godownload.jpg", downloadHttpHandler)
	//http.ListenAndServe("127.0.0.1:8000", nil)
	error := http.ListenAndServeTLS("127.0.0.1:8085", certPath, keyPath, nil)
	if error != nil {
		fmt.Printf(error.Error())
	}
}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Result struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []People `json:"data"`
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := People{
		Name: "张三",
		Age:  14,
	}
	p1 := People{
		Name: "李四",
		Age:  15,
	}
	arr := []People{p, p1}
	result := Result{
		Code:    400,
		Message: "success",
		Data:    arr,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(result)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func getHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := People{
		Name: "张三",
		Age:  14,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	contentType := r.Header.Get("Content-Type")
	fmt.Println("contentType", contentType)
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Printf("read body err, %v\n", err)
	// 	return
	// }
	// println("json:", string(body))
	reader, _ := r.MultipartReader()
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())

		if part.FileName() == "" { //formdata
			data, _ := ioutil.ReadAll(part)
			fmt.Printf("FormData=[%s]\n", string(data))
		} else {
			file, err := os.Create("./" + part.FileName())
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			io.Copy(file, part)
		}
	}
	p := People{
		Name: "张三",
		Age:  14,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func headHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := People{
		Name: "张三",
		Age:  14,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func putHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := People{
		Name: "张三",
		Age:  14,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func deleteHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := People{
		Name: "张三",
		Age:  14,
	}
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func pbHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	p := &pbout.Person{}
	p.Id = 111
	p.Name = "张三"
	w.Header().Set("Content-Type", "application/octet-stream")
	msg, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("解析错误")
	}
	w.Write(msg)
}

func downloadHttpHandler(w http.ResponseWriter, r *http.Request) {
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
	file, e := os.Open("./ic_tel.png")
	if e != nil {
		panic(e)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return
	}
	w.Write(bytes)
}
