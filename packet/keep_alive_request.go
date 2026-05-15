package packet

import "encoding/json"

type KeepAliveRequest struct {
	Seq         int `json:"-"`
	Interactive bool
}

const KeepAliveOpcode = 1

func (pk *KeepAliveRequest) MarshalPacket() ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     requestCmd,
		Seq:     pk.Seq,
		Opcode:  KeepAliveOpcode,
		Payload: pk,
	})
}

func (*KeepAliveRequest) Opcode() int {
	return KeepAliveOpcode
}
