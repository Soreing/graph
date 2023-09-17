package resource

// microsoft.graph.mediaStream
// https://learn.microsoft.com/en-us/graph/api/resources/mediastream?view=graph-rest-1.0
type MediaStream struct {
	ODataType   string `json:"@odata.type"`
	Direction   string `json:"direction,omitempty"`
	Label       string `json:"label,omitempty"`
	MediaType   string `json:"mediaType,omitempty"`
	ServerMuted *bool  `json:"serverMuted,omitempty"`
	SourceId    string `json:"sourceId,omitempty"`
}

func (r *MediaStream) GetType() (ResourceType, bool) {
	return MediaStreamResourceType, false
}

func (r *MediaStream) Validate() bool {
	return r.ODataType == MediaStreamODataType
}
