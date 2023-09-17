package resource

// microsoft.graph.participant
// https://learn.microsoft.com/en-us/graph/api/resources/participant?view=graph-rest-1.0
type Participant struct {
	ODataType            string           `json:"@odata.type"`
	Id                   string           `json:"id,omitempty"`
	Info                 *ParticipantInfo `json:"info,omitempty"`
	IsIdentityAnonymized bool             `json:"isIdentityAnonymized,omitempty"`
	IsInLobby            bool             `json:"isInLobby,omitempty"`
	IsMuted              bool             `json:"isMuted,omitempty"`
	MediaStreams         []MediaStream    `json:"mediaStreams,omitempty"`
	MeetingRole          string           `json:"meetingRole,omitempty"`
	Metadata             string           `json:"metadata,omitempty"`
	PublishedStates      []PublishedState `json:"publishedStates,omitempty"`
	RecordingInfo        *RecordingInfo   `json:"recordingInfo,omitempty"`
	ReplacementLink      string           `json:"replacementLink,omitempty"`
}

func (r *Participant) GetType() (ResourceType, bool) {
	return ParticipantInfoResourceType, false
}

func (r *Participant) Validate() bool {
	if r.ODataType != ParticipantODataType {
		return false
	}
	if r.Info != nil && !r.Info.Identity.Validate() {
		return false
	}
	if r.RecordingInfo != nil && !r.RecordingInfo.Validate() {
		return false
	}
	for _, rs := range r.MediaStreams {
		if !rs.Validate() {
			return false
		}
	}
	for _, rs := range r.PublishedStates {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
