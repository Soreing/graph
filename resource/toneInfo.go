package resource

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
