package resource

// microsoft.graph.commsNotification
// https://learn.microsoft.com/en-us/graph/api/resources/commsnotification?view=graph-rest-1.0
type CommsNotification struct {
	ODataType    string       `json:"@odata.type"`
	ChangeType   string       `json:"changeType,omitempty"`
	ResourceUrl  string       `json:"resourceUrl,omitempty"`
	ResourceData *AnyResource `json:"resourceData,omitempty"`
}

func (r *CommsNotification) GetType() (ResourceType, bool) {
	return CommsNotificationResourceType, false
}

func (r *CommsNotification) Validate() bool {
	if r.ODataType != CommsNotificationODataType {
		return false
	}
	if r.ResourceData != nil && !r.ResourceData.Validate() {
		return false
	}
	return true
}
