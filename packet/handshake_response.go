package packet

// HandshakeResponse is sent only by server after client Handshake request.
type HandshakeResponse struct {
	onlyResponse

	PhoneAuthEnabled bool     `json:"phone-auth-enabled"`
	Location         string   `json:"location"`
	Lang             bool     `json:"lang"`
	RegCountryCode   []string `json:"reg-country-code"`
}

func (pk *HandshakeResponse) Opcode() int {
	return HandshakeOpcode
}
