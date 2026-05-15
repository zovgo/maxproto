package packet

type KeepAliveResponse struct {
	onlyResponse
}

func (*KeepAliveResponse) Opcode() int {
	return KeepAliveOpcode
}
