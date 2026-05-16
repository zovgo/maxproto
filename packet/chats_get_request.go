package packet

import "encoding/json"

// ChatsGetRequest is only sent by client to request full chat details for
// specific IDs.
type ChatsGetRequest struct {
	ChatIDs []int64 `json:"chatIds"`
}

const ChatsGetOpcode = 48

func (pk *ChatsGetRequest) MarshalPacket(seq int) ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     requestCmd,
		Seq:     seq,
		Opcode:  ChatsGetOpcode,
		Payload: pk,
	})
}

func (pk *ChatsGetRequest) Opcode() int {
	return ChatsGetOpcode
}
