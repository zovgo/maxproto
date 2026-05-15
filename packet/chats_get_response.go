package packet

import "github.com/zovgo/maxproto/protocol"

// ChatsGetResponse is only sent by server after client ChatsGetRequest
type ChatsGetResponse struct {
	onlyResponse

	Chats []protocol.Chat `json:"chats"`
}

func (pk *ChatsGetResponse) Opcode() int {
	return ChatsGetOpcode
}
