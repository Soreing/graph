package resource

// microsoft.graph.call
// https://learn.microsoft.com/en-us/graph/api/resources/call?view=graph-rest-1.0
type Call struct {
	ODataType           string                      `json:"@odata.type"`
	CallbackUri         string                      `json:"callbackUri,omitempty"`
	CallChainId         string                      `json:"callChainId,omitempty"`
	CallOptions         *OutgoingCallOptions        `json:"callOptions,omitempty"`
	CallRoutes          []CallRoute                 `json:"callRoutes,omitempty"`
	ChatInfo            *ChatInfo                   `json:"chatInfo,omitempty"`
	Direction           string                      `json:"direction,omitempty"`
	Id                  string                      `json:"id,omitempty"`
	IncomingContext     *IncomingContext            `json:"incomingContext,omitempty"`
	MediaConfig         *AnyResource                `json:"mediaConfig,omitempty"`
	MediaState          string                      `json:"mediaState,omitempty"`
	MeetingInfo         *AnyResource                `json:"meetingInfo,omitempty"`
	MyParticipantId     string                      `json:"myParticipantId,omitempty"`
	RequestedModalities []string                    `json:"requestedModalities,omitempty"`
	ResultInfo          *ResultInfo                 `json:"resultInfo,omitempty"`
	Source              *ParticipantInfo            `json:"source,omitempty"`
	State               string                      `json:"state,omitempty"`
	Subject             string                      `json:"subject,omitempty"`
	Targets             []InvitationParticipantInfo `json:"targets,omitempty"`
	ToneInfo            *ToneInfo                   `json:"toneInfo,omitempty"`
	Transcription       *CallTranscriptionInfo      `json:"transcription,omitempty"`
}

func (r *Call) GetType() (ResourceType, bool) {
	return CallResourceType, false
}

func (r *Call) Validate() bool {
	if r.ODataType != CallODataType {
		return false
	}
	if r.CallOptions != nil && !r.CallOptions.Validate() {
		return false
	}
	if r.ChatInfo != nil && !r.ChatInfo.Validate() {
		return false
	}
	if r.IncomingContext != nil && !r.IncomingContext.Validate() {
		return false
	}
	if r.MediaConfig != nil && !r.MediaConfig.Validate() {
		return false
	}
	if r.MeetingInfo != nil && !r.MeetingInfo.Validate() {
		return false
	}
	if r.ResultInfo != nil && !r.ResultInfo.Validate() {
		return false
	}
	if r.Source != nil && !r.Source.Validate() {
		return false
	}
	if r.ToneInfo != nil && !r.ToneInfo.Validate() {
		return false
	}
	if r.Transcription != nil && !r.Transcription.Validate() {
		return false
	}
	for _, rs := range r.CallRoutes {
		if !rs.Validate() {
			return false
		}
	}
	for _, rs := range r.Targets {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
