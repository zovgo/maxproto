package protocol

type Keyboard struct {
	Buttons [][]KeyboardButton `json:"buttons"`
}

type KeyboardButton struct {
	Type    string `json:"type"`
	Url     string `json:"url,omitempty"`
	Text    string `json:"text"`
	Payload string `json:"payload,omitempty"`
	Intent  string `json:"intent,omitempty"`
}
