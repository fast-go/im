package logic

import (
    "fmt"
    "github.com/fast-go/websocket"
    "strconv"
)

func Online(s *websocket.Subject) {

}
// {"event":"event_message_text","data":{"to_user_id":1,"text":"您好呀"}}
func MessageText(s *websocket.Subject) {
    if  data, ok := s.MessageFormat.Data.(map[string]interface{});ok {
        toUserId,toUserIdOk := data["to_user_id"].(float64)
        text,textOk := data["text"].(string)
        if  toUserIdOk && textOk {
            fmt.Println(s.SendToUid(websocket.UniqueIdentification(strconv.Itoa(int(toUserId))),[]byte(text)))
        }
    }
}



