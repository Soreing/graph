package resource

// microsoft.graph.appHostedMediaConfig
// https://learn.microsoft.com/en-us/graph/api/resources/apphostedmediaconfig?view=graph-rest-1.0
type AppHostedMediaConfig struct {
	ODataType string `json:"@odata.type"`
	Blob      string `json:"blob,omitempty"`
}

func (r *AppHostedMediaConfig) GetType() (ResourceType, bool) {
	return AppHostedMediaConfigResourceType, false
}

func (r *AppHostedMediaConfig) Validate() bool {
	return r.ODataType == AppHostedMediaConfigODataType
}
