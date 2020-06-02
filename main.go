package im

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
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		im.Middleware(writer,request,&logic.Chain{})
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Print(err)
	}
}