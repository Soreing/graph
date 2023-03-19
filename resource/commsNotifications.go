package resource

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
	} else if len(r.Value) > 0 {
		for i := range r.Value {
			if !r.Value[i].Validate() {
				return false
			}
		}
	}
	return true
}
