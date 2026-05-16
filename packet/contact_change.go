package packet

import "github.com/zovgo/maxproto/protocol"

type ContactChange struct { //этот пакет отправляется как-то странно
	onlyResponse

	Contact protocol.Contact
}

const ContactChangeOpcode = 131

func (pk *ContactChange) Opcode() int {
	return ContactChangeOpcode
}
