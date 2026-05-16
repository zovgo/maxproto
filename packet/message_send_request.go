package packet

import (
	"encoding/json"

	"github.com/zovgo/maxproto/protocol"
)

type MessageSendRequest struct {
	ChatID  int64            `json:"chatId"`
	Message protocol.Message `json:"message"`
	Notify  bool             `json:"notify"`
}

func (pk *MessageSendRequest) MarshalPacket(seq int) ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     requestCmd,
		Seq:     seq,
		Opcode:  MessageSendOpcode,
		Payload: pk,
	})
}

const MessageSendOpcode = 64

func (*MessageSendRequest) Opcode() int {
	return MessageSendOpcode
}
