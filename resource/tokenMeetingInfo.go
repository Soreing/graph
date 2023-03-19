package resource

type TokenMeetingInfo struct {
	ODataType string `json:"@odata.type"`
	Token     string `json:"token,omitempty"`
}

func (r *TokenMeetingInfo) GetType() (ResourceType, bool) {
	return TokenMeetingInfoResourceType, false
}

func (r *TokenMeetingInfo) Validate() bool {
	return r.ODataType == TokenMeetingInfoODataType
}
