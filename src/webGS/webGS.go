package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:80", nil))
}

//处理html请求
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if r.URL.Path == "/" {
		r.URL.Path += "index.html"
	}
	str, err := getFile("../web" + r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "505")
	}
	fmt.Fprint(w, str)
}

//读取文件
func getFile(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	return string(str), err
}
