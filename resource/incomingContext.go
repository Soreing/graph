package resource

// microsoft.graph.incomingContext
// https://learn.microsoft.com/en-US/graph/api/resources/incomingcontext?view=graph-rest-beta
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
	if r.ODataType != IncomingContextODataType {
		return false
	}
	if r.OnBehalfOf != nil && !r.OnBehalfOf.Validate() {
		return false
	}
	if r.Transferor != nil && !r.Transferor.Validate() {
		return false
	}
	return true
}
