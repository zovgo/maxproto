package packet

import "encoding/json"

type KeepAliveRequest struct {
	Interactive bool
}

const KeepAliveOpcode = 1

func (pk *KeepAliveRequest) MarshalPacket(seq int) ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     requestCmd,
		Seq:     seq,
		Opcode:  KeepAliveOpcode,
		Payload: pk,
	})
}

func (*KeepAliveRequest) Opcode() int {
	return KeepAliveOpcode
}
