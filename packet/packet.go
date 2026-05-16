package packet

import "github.com/zovgo/maxproto/internal"

type Packet interface {
	MarshalPacket(seq int) ([]byte, error)
	Opcode() int
}

var pool = make(map[int]packetEntry)

func init() {
	register(newEntry(func() Packet { return new(HandshakeResponse) }))
	register(newEntry(func() Packet { return new(LoginResponse) }))
	register(newEntry(func() Packet { return new(BannersConfigResponse) }))
	register(newEntry(func() Packet { return new(ChatsGetResponse) }))
	register(newEntry(func() Packet { return new(ReceiveMessage) }))
	register(newEntry(func() Packet { return new(KeepAliveResponse) }, unmarshalKeepAliveResp))
	register(newEntry(func() Packet { return new(MessageSendResponse) }))
}

type packetEntry struct {
	opcode int
	nr     func() Packet                // new response model
	urp    func([]byte) (Packet, error) // optional unmarshal response payload func
}

func register(e packetEntry) {
	pool[e.opcode] = e
}

func newEntry(nr func() Packet, urp ...func([]byte) (Packet, error)) packetEntry {
	return packetEntry{
		opcode: nr().Opcode(),
		nr:     nr,
		urp:    internal.MustFirstOptionOr(urp, nil),
	}
}
