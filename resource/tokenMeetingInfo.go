package resource

// microsoft.graph.tokenMeetingInfo
// https://learn.microsoft.com/en-us/graph/api/resources/tokenmeetinginfo?view=graph-rest-1.0
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
