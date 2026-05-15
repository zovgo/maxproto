package protocol

type Chat struct {
	ParticipantsCount        int                 `json:"participantsCount,omitempty"`
	Access                   string              `json:"access,omitempty"`
	InvitedBy                int                 `json:"invitedBy,omitempty"`
	Type                     string              `json:"type"`
	Title                    string              `json:"title,omitempty"`
	LastFireDelayedErrorTime int64               `json:"lastFireDelayedErrorTime"`
	LastDelayedUpdateTime    int64               `json:"lastDelayedUpdateTime"`
	NewMessages              int                 `json:"newMessages,omitempty"`
	VideoConversation        *VideoConversation  `json:"videoConversation,omitempty"`
	Options                  *ChatOptions        `json:"options,omitempty"`
	Modified                 int64               `json:"modified"`
	ID                       int64               `json:"id"`
	AdminParticipants        map[int64]AdminInfo `json:"adminParticipants,omitempty"`
	Participants             map[int64]int64     `json:"participants"`
	Owner                    int64               `json:"owner"`
	JoinTime                 int64               `json:"joinTime"`
	Created                  int64               `json:"created"`
	LastMessage              *Message            `json:"lastMessage"`
	LastEventTime            int64               `json:"lastEventTime"`
	Reactions                *ChatReactions      `json:"reactions,omitempty"`
	MessagesCount            int                 `json:"messagesCount,omitempty"`
	Admins                   []int               `json:"admins,omitempty"`
	Status                   string              `json:"status"`
	CID                      int64               `json:"cid,omitempty"`
	Subject                  *ChatSubject        `json:"subject,omitempty"`
	Link                     string              `json:"link,omitempty"`
	Description              string              `json:"description,omitempty"`
	PinnedMessage            *Message            `json:"pinnedMessage,omitempty"`
	BaseRawIconURL           string              `json:"baseRawIconUrl,omitempty"`
	Restrictions             int                 `json:"restrictions,omitempty"`
	BaseIconURL              string              `json:"baseIconUrl,omitempty"`
	PrevMessageID            string              `json:"prevMessageId,omitempty"`
	HasBots                  bool                `json:"hasBots,omitempty"`
}

type AdminInfo struct {
	ID int64 `json:"id"`
}

type ChatSubject struct {
	OrganizationIds []int `json:"organizationIds"`
}

type ChatReactions struct {
	IsActive   bool  `json:"isActive"`
	UpdateTime int64 `json:"updateTime"`
}

type ChatOptions struct {
	Official                 bool `json:"OFFICIAL,omitempty"`
	OnlyOwnerCanChangeIcon   bool `json:"ONLY_OWNER_CAN_CHANGE_ICON_TITLE,omitempty"`
	OnlyAdminCanCall         bool `json:"ONLY_ADMIN_CAN_CALL,omitempty"`
	Comments                 bool `json:"COMMENTS,omitempty"`
	SentByPhone              bool `json:"SENT_BY_PHONE,omitempty"`
	ConfirmBeforeSend        bool `json:"CONFIRM_BEFORE_SEND,omitempty"`
	APlusChannel             bool `json:"A_PLUS_CHANNEL,omitempty"`
	AllCanPinMessage         bool `json:"ALL_CAN_PIN_MESSAGE,omitempty"`
	SignAdmin                bool `json:"SIGN_ADMIN,omitempty"`
	MessageCopyNotAllowed    bool `json:"MESSAGE_COPY_NOT_ALLOWED,omitempty"`
	OnlyAdminCanAddMember    bool `json:"ONLY_ADMIN_CAN_ADD_MEMBER,omitempty"`
	MembersCanSeePrivateLink bool `json:"MEMBERS_CAN_SEE_PRIVATE_LINK,omitempty"`
	ServiceChat              bool `json:"SERVICE_CHAT,omitempty"`
	ContentLevelChat         bool `json:"CONTENT_LEVEL_CHAT,omitempty"`
}
