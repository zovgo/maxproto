package protocol

type Message struct {
	Sender     int64            `json:"sender,omitempty"`
	ID         string           `json:"id"`
	Time       int64            `json:"time"`
	Text       string           `json:"text"`
	Type       string           `json:"type"`
	Attaches   []Attach         `json:"attaches"`
	Options    int              `json:"options,omitempty"`
	Cid        int64            `json:"cid,omitempty"`
	Link       *Link            `json:"link,omitempty"`
	Elements   []MessageElement `json:"elements,omitempty"`
	UpdateTime int64            `json:"updateTime,omitempty"`
}

type MessageElement struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
	From   int    `json:"from,omitempty"`
}

type Link struct {
	Type    string   `json:"type"`
	Message *Message `json:"message"`
	ChatId  int      `json:"chatId"`
}

type Attach struct {
	Type                string    `json:"_type"`
	Event               string    `json:"event,omitempty"`
	PreviewData         string    `json:"previewData,omitempty"`
	BaseUrl             string    `json:"baseUrl,omitempty"`
	PhotoToken          string    `json:"photoToken,omitempty"`
	Width               int       `json:"width,omitempty"`
	PhotoId             int64     `json:"photoId,omitempty"`
	Height              int       `json:"height,omitempty"`
	Duration            int       `json:"duration,omitempty"`
	AudioId             int64     `json:"audioId,omitempty"`
	Wave                string    `json:"wave,omitempty"`
	Url                 string    `json:"url,omitempty"`
	Token               string    `json:"token,omitempty"`
	ConversationId      string    `json:"conversationId,omitempty"`
	HangupType          string    `json:"hangupType,omitempty"`
	CallType            string    `json:"callType,omitempty"`
	ContactIds          []int     `json:"contactIds,omitempty"`
	Keyboard            *Keyboard `json:"keyboard,omitempty"`
	CallbackId          string    `json:"callbackId,omitempty"`
	JoinLink            string    `json:"joinLink,omitempty"`
	AuthorType          string    `json:"authorType,omitempty"`
	LottieUrl           string    `json:"lottieUrl,omitempty"`
	StickerId           int       `json:"stickerId,omitempty"`
	Tags                []string  `json:"tags,omitempty"`
	SetId               int       `json:"setId,omitempty"`
	Time                int64     `json:"time,omitempty"`
	StickerType         string    `json:"stickerType,omitempty"`
	Audio               bool      `json:"audio,omitempty"`
	UserIds             []int     `json:"userIds,omitempty"`
	PinnedMessage       *Message  `json:"pinnedMessage,omitempty"`
	TranscriptionStatus string    `json:"transcriptionStatus,omitempty"`
}
