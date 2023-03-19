package resource

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
	return true // TODO
}
