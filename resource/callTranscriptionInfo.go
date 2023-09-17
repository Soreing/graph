package resource

import (
	"time"
)

// microsoft.graph.callTranscriptionInfo
// https://learn.microsoft.com/en-us/graph/api/resources/calltranscriptioninfo?view=graph-rest-1.0
type CallTranscriptionInfo struct {
	ODataType            string    `json:"@odata.type"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	State                string    `json:"state,omitempty"`
}

func (r *CallTranscriptionInfo) GetType() (ResourceType, bool) {
	return CallTranscriptionInfoResourceType, false
}

func (r *CallTranscriptionInfo) Validate() bool {
	return r.ODataType == CallTranscriptionInfoODataType
}
