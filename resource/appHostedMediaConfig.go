package resource

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
