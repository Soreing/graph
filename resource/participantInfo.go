package resource

type ParticipantInfo struct {
	ODataType     string       `json:"@odata.type"`
	CountryCode   string       `json:"countryCode,omitempty"`
	EndpointType  string       `json:"endpointType,omitempty"`
	Identity      *IdentitySet `json:"identity,omitempty"`
	LanguageId    string       `json:"languageId,omitempty"`
	ParticipantId string       `json:"participantId,omitempty"`
	Region        string       `json:"region,omitempty"`
}

func (r *ParticipantInfo) GetType() (ResourceType, bool) {
	return ParticipantInfoResourceType, false
}

func (r *ParticipantInfo) Validate() bool {
	if r.ODataType != ParticipantInfoODataType {
		return false
	} else if r.Identity != nil && r.Identity.Validate() {
		return false
	}
	return true
}
