package resource

// microsoft.graph.mediaStream
// https://learn.microsoft.com/en-us/graph/api/resources/mediastream?view=graph-rest-1.0
type MediaPrompt struct {
	ODataType string     `json:"@odata.type"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
}

func (r *MediaPrompt) GetType() (ResourceType, bool) {
	return MediaPromptResourceType, false
}

func (r *MediaPrompt) Validate() bool {
	if r.ODataType != MediaPromptODataType {
		return false
	}
	if r.MediaInfo != nil && !r.MediaInfo.Validate() {
		return false
	}
	return true
}
