package packet

import "errors"

type KeepAliveResponse struct {
	onlyResponse
}

func (*KeepAliveResponse) Opcode() int {
	return KeepAliveOpcode
}

var ErrPayloadMustBeNilForKeepAlive = errors.New("payload must be nil for keep alive packet")

func unmarshalKeepAliveResp(d []byte) (Packet, error) {
	if d != nil && string(d) != "null" {
		return nil, ErrPayloadMustBeNilForKeepAlive
	}
	return new(KeepAliveResponse), nil
}
