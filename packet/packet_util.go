package packet

import "errors"

var ErrPacketDoesntSupportMarshalling = errors.New("packet does not support marshalling")

type onlyResponse struct{}

func (*onlyResponse) MarshalPacket() ([]byte, error) {
	return nil, ErrPacketDoesntSupportMarshalling
}
