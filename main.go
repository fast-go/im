package main

import (
	"fmt"
	"github.com/fast-go/websocket"
	"im/define"
	"im/logic"
	"net/http"
)


func main() {
	im := websocket.NewWebSocket()
	im.Events.Register(define.EVENT_ONLINE, logic.Online)
	im.Events.Register(define.EVENT_MESSAGE_TEXT, logic.MessageText)

	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		im.Middleware(writer,request,&logic.Chain{})
	})
	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(11)
}