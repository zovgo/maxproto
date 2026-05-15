package packet

import (
	"encoding/json"

	"github.com/zovgo/maxproto/protocol"
)

// LoginResponse is only sent from server after client LoginRequest.
type LoginResponse struct {
	onlyResponse

	VideoChatHistory bool               `json:"videoChatHistory"`
	Profile          protocol.Profile   `json:"profile"`
	Chats            []protocol.Chat    `json:"chats"`
	ChatMarker       int64              `json:"chatMarker"`
	Messages         json.RawMessage    `json:"messages"` //TODO
	Time             int64              `json:"time"`
	Presence         json.RawMessage    `json:"presence"` //TODO
	Config           protocol.Config    `json:"config"`
	Contacts         []protocol.Contact `json:"contacts"`
}

func (pk *LoginResponse) Opcode() int {
	return LoginOpcode
}
