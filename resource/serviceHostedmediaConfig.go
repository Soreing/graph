package resource

type ServiceHostedMediaConfig struct {
	ODataType     string      `json:"@odata.type"`
	PreFetchMedia []MediaInfo `json:"preFetchMedia,omitempty"`
}

func (r *ServiceHostedMediaConfig) GetType() (ResourceType, bool) {
	return ServiceHostedMediaConfigResourceType, false
}

func (r *ServiceHostedMediaConfig) Validate() bool {
	if r.ODataType != ServiceHostedMediaConfigODataType {
		return false
	} else if len(r.PreFetchMedia) > 0 {
		for i := range r.PreFetchMedia {
			if !r.PreFetchMedia[i].Validate() {
				return false
			}
		}
	}
	return true
}
