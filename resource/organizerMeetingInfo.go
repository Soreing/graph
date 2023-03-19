package resource

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
	} else if r.Organizer != nil && !r.Organizer.Validate() {
		return false
	}
	return true
}
