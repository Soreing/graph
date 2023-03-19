package resource

import (
	"time"
)

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
