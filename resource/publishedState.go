package resource

// TODO: Find documentation
type PublishedState struct {
	ODataType string `json:"@odata.type"`
	TypeRank  int    `json:"typeRank,omitempty"`
	Type      string `json:"type,omitempty"`
	StateId   string `json:"stateId,omitempty"`
}

func (r *PublishedState) GetType() (ResourceType, bool) {
	return PublishedStateResourceType, false
}

func (r *PublishedState) Validate() bool {
	return r.ODataType == PublishedStateODataType
}
