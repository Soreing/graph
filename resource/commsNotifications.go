package resource

// microsoft.graph.commsNotifications
// https://learn.microsoft.com/en-us/graph/api/resources/commsnotifications?view=graph-rest-1.0
type CommsNotifications struct {
	ODataType string              `json:"@odata.type"`
	Value     []CommsNotification `json:"value,omitempty"`
}

func (r *CommsNotifications) GetType() (ResourceType, bool) {
	return CommsNotificationsResourceType, false
}

func (r *CommsNotifications) Validate() bool {
	if r.ODataType != CommsNotificationsODataType {
		return false
	}
	for _, rs := range r.Value {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
