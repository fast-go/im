package define

import "github.com/fast-go/websocket"

const (
	EVENT_ONLINE = "event_online"
	EVENT_ENTER = "event_enter"
	EVENT_LAYOUT = "event_layout"
	EVENT_MESSAGE_TEXT = "event_message_text"
	EVENT_MESSAGE_IMG = "event_message_img"
	EVENT_MESSAGE_VIDEO = "event_message_video"
)

type OnlineRequest struct {

}

type OnlineResponse struct {

}

type MessageTextReceive struct {
	ToUserId websocket.UniqueIdentification `json:"to_user_id"`
	Text  string `json:"text"`
}

type MessageTextSending struct {

}


