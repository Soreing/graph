package resource

// microsoft.graph.joinMeetingIdMeetingInfo
// https://learn.microsoft.com/en-us/graph/api/resources/joinmeetingidmeetinginfo?view=graph-rest-1.0
type JoinMeetingIdMeetingInfo struct {
	ODataType     string `json:"@odata.type"`
	JoinMeetingId string `json:"joinMeetingId,omitempty"`
	Passcode      string `json:"passcode,omitempty"`
}

func (r *JoinMeetingIdMeetingInfo) GetType() (ResourceType, bool) {
	return JoinMeetingIdMeetingInfoResourceType, false
}

func (r *JoinMeetingIdMeetingInfo) Validate() bool {
	return r.ODataType == JoinMeetingIdMeetingInfoODataType
}
