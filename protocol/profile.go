package protocol

type Profile struct {
	ProfileOptions []any   `json:"profileOptions"`
	Contact        Contact `json:"contact"`
}
