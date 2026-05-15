package packet

import "encoding/json"

type KeepAliveRequest struct {
	Seq         int
	Interactive bool
}

const KeepAliveOpcode = 1

const KeepAliveRequestCmd = 0

func (pk *KeepAliveRequest) MarshalPacket() ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     KeepAliveRequestCmd,
		Seq:     pk.Seq,
		Opcode:  KeepAliveOpcode,
		Payload: pk,
	})
}

func (*KeepAliveRequest) Opcode() int {
	return KeepAliveOpcode
}
