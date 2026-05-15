package protocol

import "encoding/json"

type VideoConversation struct {
	JoinLink              string            `json:"joinLink"`
	Type                  int               `json:"type,omitempty"`
	PreviewParticipantIDs []json.RawMessage `json:"previewParticipantIds"`
	ConversationID        string            `json:"conversationId"`
	CallType              string            `json:"callType"`
}
