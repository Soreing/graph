package resource

type Call struct {
	ODataType           string                 `json:"@odata.type"`
	CallbackUri         string                 `json:"callbackUri,omitempty"`
	CallChainId         string                 `json:"callChainId,omitempty"`
	CallOptions         *OutgoingCallOptions   `json:"callOptions,omitempty"`
	CallRoutes          []CallRoute            `json:"callRoutes,omitempty"`
	ChatInfo            *ChatInfo              `json:"chatInfo,omitempty"`
	Direction           string                 `json:"direction,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	IncomingContext     *IncomingContext       `json:"incomingContext,omitempty"`
	MediaConfig         *AnyResource           `json:"mediaConfig,omitempty"`
	MediaState          string                 `json:"mediaState,omitempty"`
	MeetingInfo         *AnyResource           `json:"meetingInfo,omitempty"`
	MyParticipantId     string                 `json:"myParticipantId,omitempty"`
	RequestedModalities []string               `json:"requestedModalities,omitempty"`
	ResultInfo          *ResultInfo            `json:"resultInfo,omitempty"`
	Source              *ParticipantInfo       `json:"source,omitempty"`
	State               string                 `json:"state,omitempty"`
	Subject             string                 `json:"subject,omitempty"`
	Targets             []ParticipantInfo      `json:"targets,omitempty"`
	ToneInfo            *ToneInfo              `json:"toneInfo,omitempty"`
	Transcription       *CallTranscriptionInfo `json:"transcription,omitempty"`
}

func (r *Call) GetType() (ResourceType, bool) {
	return CallResourceType, false
}

func (r *Call) Validate() bool {
	// TODO: Finish...
	return true
}
