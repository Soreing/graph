package resource

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
