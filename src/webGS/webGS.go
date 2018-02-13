package main

import (
	"fmt"
	//"golang.org/x/net/websocket"
	"game"
	ws "github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/game/websocket", webSocketHandler)
	log.Fatal(http.ListenAndServe("localhost:80", nil))
}

//处理html请求
func httpHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if r.URL.Path == "/" {
		r.URL.Path += "index.html"
	}
	str, err := getFile("./web" + r.URL.Path) //读取文件
	if err != nil {
		fmt.Fprintf(w, "找不到文件")
	}
	fmt.Fprint(w, str)
}

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//处理WebSocket请求
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go game.ServePlayer(game.CreatePlayer(conn)) //服务玩家
}

//读取文件
func getFile(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	return string(str), err
}
