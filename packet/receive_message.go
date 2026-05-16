package packet

import "github.com/zovgo/maxproto/protocol"

// fixme: client sends this ? why
//  {"ver":11,"cmd":1,"seq":22,"opcode":128,"payload":{"chatId":-7777,"messageId":"1111"}}

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
