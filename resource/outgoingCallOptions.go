package resource

// microsoft.graph.outgoingCallOptions
// https://learn.microsoft.com/en-us/graph/api/resources/outgoingcalloptions?view=graph-rest-1.0
type OutgoingCallOptions struct {
	ODataType                           string `json:"@odata.type"`
	HideBotAfterEscalation              *bool  `json:"hideBotAfterEscalation,omitempty"`
	IsContentSharingNotificationEnabled *bool  `json:"isContentSharingNotificationEnabled,omitempty"`
}

func (r *OutgoingCallOptions) GetType() (ResourceType, bool) {
	return OutgoingCallOptionsResourceType, false
}

func (r *OutgoingCallOptions) Validate() bool {
	return r.ODataType == OutgoingCallOptionsODataType
}
