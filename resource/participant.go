package resource

type Participant struct {
	ODataType     string           `json:"@odata.type"`
	Id            string           `json:"id,omitempty"`
	Info          *ParticipantInfo `json:"info,omitempty"`
	IsInLobby     bool             `json:"isInLobby,omitempty"`
	IsMuted       bool             `json:"isMuted,omitempty"`
	MediaStreams  []MediaStream    `json:"mediaStreams,omitempty"`
	Metadata      string           `json:"metadata,omitempty"`
	RecordingInfo *RecordingInfo   `json:"recordingInfo,omitempty"`
}

func (r *Participant) GetType() (ResourceType, bool) {
	return ParticipantInfoResourceType, false
}

func (r *Participant) Validate() bool {
	if r.ODataType != ParticipantODataType {
		return false
	} else if r.Info != nil && !r.Info.Identity.Validate() {
		return false
	} else if r.RecordingInfo != nil && !r.RecordingInfo.Initiator.Validate() {
		return false
	} else if len(r.MediaStreams) > 0 {
		for i := range r.MediaStreams {
			if !r.MediaStreams[i].Validate() {
				return false
			}
		}
	}
	return true

}
