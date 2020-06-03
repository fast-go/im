package logic

import (
	"github.com/fast-go/websocket"
	"net/http"
	"strconv"
	"time"
)

type Chain struct {

}


var UserIds []websocket.UniqueIdentification

func init()  {
	for i:= 1 ;i < 100 ;i ++ {
		UserIds = append(UserIds,websocket.UniqueIdentification(strconv.Itoa(i)))
	}
}

func GetUserId() websocket.UniqueIdentification  {
	uid := UserIds[0]
	UserIds = UserIds[1:]
	return uid
}

// 身份验证
func (*Chain)Identity(w http.ResponseWriter, r *http.Request) (error,websocket.UniqueIdentification){
	return nil,GetUserId()
}

// 链接之前
func (*Chain) ConnBefore(w http.ResponseWriter, r *http.Request) {

}

// 链接成功
func (*Chain)ConnDone(c *websocket.Connection) {
	//todo 可以自行存储链接状态,可对链接进行分组管理.组消息发送的时候只需要遍历制定组的链接
	_ = c.WriteMessage([]byte("您的userid:" + c.UniqueIdentification))
}

// 心跳检测
func (*Chain)Heartbeat(c *websocket.Connection) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C
		if err := c.WriteMessage([]byte("ping"));err != nil {
			return
		}
	}
}
