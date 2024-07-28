package resource

// microsoft.graph.inviteParticipantsOperation
// https://learn.microsoft.com/en-us/graph/api/resources/inviteparticipantsoperation?view=graph-rest-1.0
type InviteParticipantsOperation struct {
	ODataType     string                      `json:"@odata.type"`
	ClientContext string                      `json:"clientContext,omitempty"`
	Id            string                      `json:"id,omitempty"`
	Participants  []InvitationParticipantInfo `json:"participants,omitempty"`
	ResultInfo    *ResultInfo                 `json:"resultInfo,omitempty"`
	Status        string                      `json:"status,omitempty"`
}

func (r *InviteParticipantsOperation) GetType() (ResourceType, bool) {
	return InviteParticipantsOperationResourceType, false
}

func (r *InviteParticipantsOperation) Validate() bool {
	if r.ODataType != InviteParticipantsOperationODataType {
		return false
	}
	if r.ResultInfo != nil && !r.ResultInfo.Validate() {
		return false
	}
	for _, rs := range r.Participants {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
