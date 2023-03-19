package resource

type MediaInfo struct {
	ODataType  string `json:"@odata.type"`
	ResourceId string `json:"resourceId,omitempty"`
	Uri        string `json:"uri,omitempty"`
}

func (r *MediaInfo) GetType() (ResourceType, bool) {
	return MediaInfoResourceType, false
}

func (r *MediaInfo) Validate() bool {
	return r.ODataType == MediaInfoODataType
}
