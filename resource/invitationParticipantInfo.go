package resource

// microsoft.graph.participantInfo
// https://learn.microsoft.com/en-us/graph/api/resources/participantinfo?view=graph-rest-1.0
type InvitationParticipantInfo struct {
	ODataType       string       `json:"@odata.type"`
	ClientVersion   string       `json:"clientVersion,omitempty"`
	CountryCode     string       `json:"countryCode,omitempty"`
	EndpointId      string       `json:"endpointId,omitempty"`
	EndpointType    string       `json:"endpointType,omitempty"`
	Identity        *IdentitySet `json:"identity,omitempty"`
	LanguageId      string       `json:"languageId,omitempty"`
	ParticipantId   string       `json:"participantId,omitempty"`
	Region          string       `json:"region,omitempty"`
	ReplacementLink string       `json:"replacementLink,omitempty"`
}

func (r *InvitationParticipantInfo) GetType() (ResourceType, bool) {
	return InvitationParticipantInfoResourceType, false
}

func (r *InvitationParticipantInfo) Validate() bool {
	if r.ODataType != InvitationParticipantInfoODataType {
		return false
	}
	if r.Identity != nil && !r.Identity.Validate() {
		return false
	}
	return true
}
