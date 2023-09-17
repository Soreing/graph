package resource

// microsoft.graph.serviceHostedMediaConfig
// https://learn.microsoft.com/en-us/graph/api/resources/servicehostedmediaconfig?view=graph-rest-1.0
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
	}
	for _, rs := range r.PreFetchMedia {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
