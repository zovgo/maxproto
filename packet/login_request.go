package packet

import "encoding/json"

// LoginRequest is only sent by client after server HandshakeResponse.
type LoginRequest struct {
	Token        string `json:"token"`
	ChatCount    int    `json:"chatsCount"`
	Interactive  bool   `json:"interactive"`
	ChatSync     int    `json:"chatsSync"`
	ContactsSync int    `json:"contactsSync"`
	PresenceSync int    `json:"presenceSync"`
	DraftsSync   int    `json:"draftsSync"`
}

const LoginOpcode = 19

const (
	LoginRequestCmd = 0
	LoginRequestSeq = 1
)

func (pk *LoginRequest) MarshalPacket() ([]byte, error) {
	return json.Marshal(packet{
		Ver:     Version,
		Cmd:     LoginRequestCmd,
		Seq:     LoginRequestSeq,
		Opcode:  LoginOpcode,
		Payload: pk,
	})
}

func (pk *LoginRequest) Opcode() int {
	return LoginOpcode
}
