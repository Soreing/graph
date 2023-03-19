package resource

type IncomingContext struct {
	ODataType             string       `json:"@odata.type"`
	SourceParticipantId   string       `json:"sourceParticipantId,omitempty"`
	ObservedParticipantId string       `json:"observedParticipantId,omitempty"`
	OnBehalfOf            *IdentitySet `json:"onBehalfOf,omitempty"`
	Transferor            *IdentitySet `json:"transferor,omitempty"`
}

func (r *IncomingContext) GetType() (ResourceType, bool) {
	return IncomingContextResourceType, false
}

func (r *IncomingContext) Validate() bool {
	if r.ODataType == IncomingContextODataType {
		return false
	} else if r.OnBehalfOf != nil && !r.OnBehalfOf.Validate() {
		return false
	} else if r.Transferor != nil && !r.Transferor.Validate() {
		return false
	}
	return true
}
