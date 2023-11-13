package resource

// microsoft.graph.callMediaState
// https://learn.microsoft.com/en-us/graph/api/resources/callmediastate?view=graph-rest-1.0
type CallMediaState struct {
	ODataType string `json:"@odata.type"`
	Audio     string `json:"audio"`
}

func (r *CallMediaState) GetType() (ResourceType, bool) {
	return CallMediaStateResourceType, false
}

func (r *CallMediaState) Validate() bool {
	if r.ODataType != CallMediaStateODataType {
		return false
	}
	if r.Audio != "active" &&
		r.Audio != "inactive" &&
		r.Audio != "unknownFutureValue" {
		return false
	}
	return true
}
