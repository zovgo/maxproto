package packet

import (
	"encoding/json"
	"errors"
)

type Packet interface {
	MarshalPacket() ([]byte, error)
	Opcode() int
}

const Version = 11

const requestCmd = 0

type packet struct {
	Ver     int `json:"ver"`
	Cmd     int `json:"cmd"`
	Seq     int `json:"seq"`
	Opcode  int `json:"opcode"`
	Payload any `json:"payload"`
}

var ErrUnhandledOpcode = errors.New("unhandled opcode")

func UnmarshalResponse(data []byte) (Packet, error) {
	var header struct {
		Ver     int             `json:"ver"`
		Cmd     int             `json:"cmd"`
		Seq     int             `json:"seq"`
		Opcode  int             `json:"opcode"`
		Payload json.RawMessage `json:"payload"`
	}
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}
	switch header.Opcode { //fixme: that's pretty much temporary solution, we need to reg packets
	case HandshakeOpcode:
		resp := new(HandshakeResponse)
		return unmarshalPayload(header.Payload, resp)
	case LoginOpcode:
		resp := new(LoginResponse)
		return unmarshalPayload(header.Payload, resp)
	case ReceiveMessageOpcode:
		resp := new(ReceiveMessage)
		return unmarshalPayload(header.Payload, resp)
	case ChatsGetOpcode:
		resp := new(ChatsGetResponse)
		return unmarshalPayload(header.Payload, resp)
	case BannersConfigOpcode:
		resp := new(BannersConfigResponse)
		return unmarshalPayload(header.Payload, resp)
	case KeepAliveOpcode:
		if header.Payload != nil && string(header.Payload) != "null" { //why
			return nil, errors.New("payload must be nil for keep alive packet")
		}
		return new(KeepAliveResponse), nil
	}
	return nil, ErrUnhandledOpcode
}

func unmarshalPayload(data []byte, target Packet) (Packet, error) {
	if err := json.Unmarshal(data, target); err != nil {
		return nil, err
	}
	return target, nil
}
