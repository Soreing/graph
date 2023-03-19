package resource

type ResourceType int

const (
	UnknownResourceType = iota
	MixedResourceType
	AppHostedMediaConfigResourceType
	CallResourceType
	CallTranscriptionInfoResourceType
	ChatInfoResourceType
	CommsNotificationsResourceType
	CommsNotificationResourceType
	OutgoingCallOptionsResourceType
	IdentitySetResourceType
	IdentityResourceType
	IncomingContextResourceType
	JoinMeetingIdMeetingInfoResourceType
	MediaInfoResourceType
	MediaStreamResourceType
	OrganizerMeetingInfoResourceType
	ParticipantResourceType
	ParticipantInfoResourceType
	RecordingInfoResourceType
	ResultInfoResourceType
	ServiceHostedMediaConfigResourceType
	TokenMeetingInfoResourceType
	ToneInfoResourceType
)

const (
	ODataType_Key = "@odata.type"

	AppHostedMediaConfigODataType     = "#microsoft.graph.appHostedMediaConfig"
	CallODataType                     = "#microsoft.graph.call"
	CallTranscriptionInfoODataType    = "#microsoft.graph.callTranscriptionInfo"
	ChatInfoODataType                 = "#microsoft.graph.chatInfo"
	CommsNotificationsODataType       = "#microsoft.graph.commsNotifications"
	CommsNotificationODataType        = "#microsoft.graph.commsNotification"
	OutgoingCallOptionsODataType      = "#microsoft.graph.outgoingCallOptions"
	IdentitySetODataType              = "#microsoft.graph.identitySet"
	IdentityODataType                 = "#microsoft.graph.identity"
	IncomingContextODataType          = "#microsoft.graph.incomingContext"
	JoinMeetingIdMeetingInfoODataType = "#microsoft.graph.joinMeetingIdMeetingInfo"
	MediaInfoODataType                = "#microsoft.graph.mediaInfo"
	MediaStreamODataType              = "#microsoft.graph.mediaStream"
	OrganizerMeetingInfoODataType     = "#microsoft.graph.organizerMeetingInfo"
	ParticipantODataType              = "#microsoft.graph.participant"
	ParticipantInfoODataType          = "#microsoft.graph.participantInfo"
	RecordingInfoODataType            = "#microsoft.graph.recordingInfo"
	ResultInfoODataType               = "#microsoft.graph.resultInfo"
	ServiceHostedMediaConfigODataType = "#microsoft.graph.serviceHostedMediaConfig"
	TokenMeetingInfoODataType         = "#microsoft.graph.tokenMeetingInfo"
	ToneInfoODataType                 = "#microsoft.graph.toneInfo"
)

func GetType(ODataType []byte) ResourceType {
	switch len(ODataType) {
	case 21:
		switch string(ODataType) {
		case CallODataType:
			return CallResourceType
		}
	case 25:
		switch string(ODataType) {
		case ChatInfoODataType:
			return ChatInfoResourceType
		case IdentityODataType:
			return IdentityResourceType
		case ToneInfoODataType:
			return ToneInfoResourceType
		}
	case 26:
		switch string(ODataType) {
		case MediaInfoODataType:
			return MediaInfoResourceType
		}
	case 28:
		switch string(ODataType) {
		case IdentitySetODataType:
			return IdentitySetResourceType
		case MediaStreamODataType:
			return MediaStreamResourceType
		case ParticipantODataType:
			return ParticipantResourceType
		}
	case 27:
		switch string(ODataType) {
		case ResultInfoODataType:
			return ResultInfoResourceType
		}
	case 30:
		switch string(ODataType) {
		case RecordingInfoODataType:
			return RecordingInfoResourceType
		}
	case 32:
		switch string(ODataType) {
		case IncomingContextODataType:
			return IncomingContextResourceType
		case ParticipantInfoODataType:
			return ParticipantInfoResourceType
		}
	case 33:
		switch string(ODataType) {
		case TokenMeetingInfoODataType:
			return TokenMeetingInfoResourceType
		}
	case 34:
		switch string(ODataType) {
		case CommsNotificationODataType:
			return CommsNotificationResourceType
		}
	case 35:
		switch string(ODataType) {
		case CommsNotificationsODataType:
			return CommsNotificationsResourceType
		}
	case 36:
		switch string(ODataType) {
		case OutgoingCallOptionsODataType:
			return OutgoingCallOptionsResourceType
		}
	case 37:
		switch string(ODataType) {
		case AppHostedMediaConfigODataType:
			return AppHostedMediaConfigResourceType
		case OrganizerMeetingInfoODataType:
			return OrganizerMeetingInfoResourceType
		}
	case 38:
		switch string(ODataType) {
		case CallTranscriptionInfoODataType:
			return CallTranscriptionInfoResourceType
		}
	case 41:
		switch string(ODataType) {
		case JoinMeetingIdMeetingInfoODataType:
			return JoinMeetingIdMeetingInfoResourceType
		case ServiceHostedMediaConfigODataType:
			return ServiceHostedMediaConfigResourceType
		}
	}
	return UnknownResourceType
}

func NewResource(ODataType []byte) Resource {
	switch len(ODataType) {
	case 21:
		switch string(ODataType) {
		case CallODataType:
			return &Call{}
		}
	case 25:
		switch string(ODataType) {
		case ChatInfoODataType:
			return &ChatInfo{}
		case IdentityODataType:
			return &Identity{}
		case ToneInfoODataType:
			return &ToneInfo{}
		}
	case 26:
		switch string(ODataType) {
		case MediaInfoODataType:
			return &MediaInfo{}
		}
	case 28:
		switch string(ODataType) {
		case IdentitySetODataType:
			return &IdentitySet{}
		case MediaStreamODataType:
			return &MediaStream{}
		case ParticipantODataType:
			return &Participant{}
		}
	case 27:
		switch string(ODataType) {
		case ResultInfoODataType:
			return &ResultInfo{}
		}
	case 30:
		switch string(ODataType) {
		case RecordingInfoODataType:
			return &RecordingInfo{}
		}
	case 32:
		switch string(ODataType) {
		case IncomingContextODataType:
			return &IncomingContext{}
		case ParticipantInfoODataType:
			return &ParticipantInfo{}
		}
	case 33:
		switch string(ODataType) {
		case TokenMeetingInfoODataType:
			return &TokenMeetingInfo{}
		}
	case 34:
		switch string(ODataType) {
		case CommsNotificationODataType:
			return &CommsNotification{}
		}
	case 35:
		switch string(ODataType) {
		case CommsNotificationsODataType:
			return &CommsNotifications{}
		}
	case 36:
		switch string(ODataType) {
		case OutgoingCallOptionsODataType:
			return &OutgoingCallOptions{}
		}
	case 37:
		switch string(ODataType) {
		case AppHostedMediaConfigODataType:
			return &AppHostedMediaConfig{}
		case OrganizerMeetingInfoODataType:
			return &OrganizerMeetingInfo{}
		}
	case 38:
		switch string(ODataType) {
		case CallTranscriptionInfoODataType:
			return &CallTranscriptionInfo{}
		}
	case 41:
		switch string(ODataType) {
		case JoinMeetingIdMeetingInfoODataType:
			return &JoinMeetingIdMeetingInfo{}
		case ServiceHostedMediaConfigODataType:
			return &ServiceHostedMediaConfig{}
		}
	}
	return nil
}
