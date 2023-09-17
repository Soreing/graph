package resource

// microsoft.graph.participantInfo
// https://learn.microsoft.com/en-us/graph/api/resources/participantinfo?view=graph-rest-1.0
type ParticipantInfo struct {
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

func (r *ParticipantInfo) GetType() (ResourceType, bool) {
	return ParticipantInfoResourceType, false
}

func (r *ParticipantInfo) Validate() bool {
	if r.ODataType != ParticipantInfoODataType {
		return false
	}
	if r.Identity != nil && !r.Identity.Validate() {
		return false
	}
	return true
}
