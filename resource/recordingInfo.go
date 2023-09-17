package resource

// microsoft.graph.recordingInfo
// https://learn.microsoft.com/en-us/graph/api/resources/recordinginfo?view=graph-rest-1.0
type RecordingInfo struct {
	ODataType       string       `json:"@odata.type"`
	Initiator       *IdentitySet `json:"initiator,omitempty"`
	RecordingStatus string       `json:"recordingStatus,omitempty"`
}

func (r *RecordingInfo) GetType() (ResourceType, bool) {
	return RecordingInfoResourceType, false
}

func (r *RecordingInfo) Validate() bool {
	if r.ODataType != RecordingInfoODataType {
		return false
	}
	if r.Initiator != nil && !r.Initiator.Validate() {
		return false
	}
	return true
}
