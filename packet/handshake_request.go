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

const (
	HandshakeRequestCmd = 0
	HandshakeRequestSeq = 0
)

func (pk *HandshakeRequest) MarshalPacket() ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     HandshakeRequestCmd,
		Seq:     HandshakeRequestSeq,
		Opcode:  HandshakeOpcode,
		Payload: pk,
	})
}

func (pk *HandshakeRequest) Opcode() int {
	return HandshakeOpcode
}
