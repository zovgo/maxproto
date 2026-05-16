package packet

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/zovgo/maxproto/protocol"
)

// HandshakeRequest is only sent by client and is the first packet in login
// sequence.
type HandshakeRequest struct {
	UserAgent *protocol.UserAgent `json:"userAgent"`
	DeviceID  uuid.UUID           `json:"deviceId"`
}

const HandshakeOpcode = 6

func (pk *HandshakeRequest) MarshalPacket(seq int) ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     requestCmd,
		Seq:     seq,
		Opcode:  HandshakeOpcode,
		Payload: pk,
	})
}

func (pk *HandshakeRequest) Opcode() int {
	return HandshakeOpcode
}
