package packet

import "github.com/zovgo/maxproto/protocol"

type ProfileChange struct { //это отправляется только для своего профиля
	onlyResponse
	Profile protocol.Profile `json:"profile"`
}

const ProfileChangeOpcode = 159

func (*ProfileChange) Opcode() int {
	return ProfileChangeOpcode
}
