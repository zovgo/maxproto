package packet

import (
	"encoding/json"
	"errors"
)

const (
	Version    = 11
	requestCmd = 0
)

type packet struct {
	Ver     int `json:"ver"`
	Cmd     int `json:"cmd"`
	Seq     int `json:"seq"`
	Opcode  int `json:"opcode"`
	Payload any `json:"payload"`
}

type header struct {
	Ver     int             `json:"ver"`
	Cmd     int             `json:"cmd"`
	Seq     int             `json:"seq"`
	Opcode  int             `json:"opcode"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

var ErrUnhandledOpcode = errors.New("unhandled opcode")

func UnmarshalResponse(data []byte) (Packet, error) {
	var h header
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	e, ok := pool[h.Opcode]
	if !ok {
		return nil, ErrUnhandledOpcode
	}
	if e.urp != nil {
		return e.urp(h.Payload)
	}
	return unmarshalPayload(e, h.Payload)
}

func unmarshalPayload(e packetEntry, payload []byte) (Packet, error) {
	pk := e.nr()
	if err := json.Unmarshal(payload, pk); err != nil {
		return nil, err
	}
	return pk, nil
}
