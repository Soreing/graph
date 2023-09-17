package resource

// microsoft.graph.chatInfo
// https://learn.microsoft.com/en-us/graph/api/resources/chatinfo?view=graph-rest-1.0
type ChatInfo struct {
	ODataType           string `json:"@odata.type"`
	MessageId           string `json:"messageId,omitempty"`
	ReplyChainMessageId string `json:"replyChainMessageId,omitempty"`
	ThreadId            string `json:"threadId,omitempty"`
}

func (r *ChatInfo) GetType() (ResourceType, bool) {
	return ChatInfoResourceType, false
}

func (r *ChatInfo) Validate() bool {
	return r.ODataType == ChatInfoODataType
}
