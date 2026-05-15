package packet

import (
	"encoding/json"

	"github.com/zovgo/maxproto/protocol"
)

// BannersConfigResponse is only sent by server in pair with LoginResponse
// (after).
type BannersConfigResponse struct {
	onlyResponse

	Banners    []json.RawMessage  `json:"banners"` //TODO
	UpdateTime protocol.Timestamp `json:"updateTime"`
	ShowTime   protocol.Timestamp `json:"showTime"`
}

const BannersConfigOpcode = 292

func (pk *BannersConfigResponse) Opcode() int {
	return BannersConfigOpcode
}
