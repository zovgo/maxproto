package packet

import "github.com/zovgo/maxproto/protocol"

type MessageSendResponse struct {
	onlyResponse

	ChatID  int64            `json:"chatId"`
	Message protocol.Message `json:"message"`
	Unread  int              `json:"unread"`
	Mark    int64            `json:"mark"`
}

func (*MessageSendResponse) Opcode() int {
	return MessageSendOpcode
}
