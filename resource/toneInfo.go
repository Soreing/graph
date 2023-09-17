package resource

// microsoft.graph.toneInfo
// https://learn.microsoft.com/en-us/graph/api/resources/toneinfo?view=graph-rest-1.0
type ToneInfo struct {
	ODataType  string `json:"@odata.type"`
	SequenceId int64  `json:"sequenceId,omitempty"`
	Tone       string `json:"tone,omitempty"`
}

func (r *ToneInfo) GetType() (ResourceType, bool) {
	return ToneInfoResourceType, false
}

func (r *ToneInfo) Validate() bool {
	return r.ODataType == ToneInfoODataType
}
