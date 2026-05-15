package packet

import "github.com/zovgo/maxproto/protocol"

type ReceiveMessage struct {
	onlyResponse

	ChatID        int64             `json:"chatId"`
	Unread        int64             `json:"unread"`
	Message       *protocol.Message `json:"message"`
	TTL           bool              `json:"ttl"`
	Mark          int64             `json:"mark"`
	PrevMessageID string            `json:"prevMessageId"`
}

const ReceiveMessageOpcode = 128

func (pk *ReceiveMessage) Opcode() int {
	return ReceiveMessageOpcode
}
