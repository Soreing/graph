package resource

// microsoft.graph.organizerMeetingInfo
// https://learn.microsoft.com/en-us/graph/api/resources/organizermeetinginfo?view=graph-rest-1.0
type OrganizerMeetingInfo struct {
	ODataType string       `json:"@odata.type"`
	Organizer *IdentitySet `json:"organizer,omitempty"`
}

func (r *OrganizerMeetingInfo) GetType() (ResourceType, bool) {
	return OrganizerMeetingInfoResourceType, false
}

func (r *OrganizerMeetingInfo) Validate() bool {
	if r.ODataType != OrganizerMeetingInfoODataType {
		return false
	}
	if r.Organizer != nil && !r.Organizer.Validate() {
		return false
	}
	return true
}
