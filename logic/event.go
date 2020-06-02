package logic

import (
    "encoding/json"
    "github.com/fast-go/websocket"
    "im/define"
)

func Online(s *websocket.Subject) {

}

func MessageText(s *websocket.Subject) {
    var data define.MessageTextReceive
    if  d, ok := s.MessageFormat.Data.(string);ok {
        if json.Unmarshal([]byte(d),&data) == nil {
            _ = s.SendToUid(data.ToUserId,[]byte(data.Text))
        }
    }
}



