package resource

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
