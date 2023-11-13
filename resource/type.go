package resource

type ResourceType int

const (
	NullResourceType = iota
	UnknownResourceType
	MixedResourceType
	AppHostedMediaConfigResourceType
	CallResourceType
	CallMediaStateResourceType
	CallRouteResourceType
	CallTranscriptionInfoResourceType
	ChatInfoResourceType
	CommsNotificationsResourceType
	CommsNotificationResourceType
	OutgoingCallOptionsResourceType
	IdentitySetResourceType
	IdentityResourceType
	IncomingContextResourceType
	InvitationParticipantInfoResourceType
	JoinMeetingIdMeetingInfoResourceType
	MediaInfoResourceType
	MediaStreamResourceType
	OrganizerMeetingInfoResourceType
	ParticipantResourceType
	ParticipantInfoResourceType
	PublishedStateResourceType
	RecordingInfoResourceType
	ResultInfoResourceType
	ServiceHostedMediaConfigResourceType
	TokenMeetingInfoResourceType
	ToneInfoResourceType
)

const (
	ODataTypeKey = "@odata.type"

	AppHostedMediaConfigODataType      = "#microsoft.graph.appHostedMediaConfig"
	CallODataType                      = "#microsoft.graph.call"
	CallMediaStateODataType            = "#microsoft.graph.callMediaState"
	CallRouteODataType                 = "#microsoft.graph.callRoute"
	CallTranscriptionInfoODataType     = "#microsoft.graph.callTranscriptionInfo"
	ChatInfoODataType                  = "#microsoft.graph.chatInfo"
	CommsNotificationsODataType        = "#microsoft.graph.commsNotifications"
	CommsNotificationODataType         = "#microsoft.graph.commsNotification"
	OutgoingCallOptionsODataType       = "#microsoft.graph.outgoingCallOptions"
	IdentitySetODataType               = "#microsoft.graph.identitySet"
	IdentityODataType                  = "#microsoft.graph.identity"
	IncomingContextODataType           = "#microsoft.graph.incomingContext"
	InvitationParticipantInfoODataType = "#microsoft.graph.invitationParticipantInfo"
	JoinMeetingIdMeetingInfoODataType  = "#microsoft.graph.joinMeetingIdMeetingInfo"
	MediaInfoODataType                 = "#microsoft.graph.mediaInfo"
	MediaStreamODataType               = "#microsoft.graph.mediaStream"
	OrganizerMeetingInfoODataType      = "#microsoft.graph.organizerMeetingInfo"
	ParticipantODataType               = "#microsoft.graph.participant"
	ParticipantInfoODataType           = "#microsoft.graph.participantInfo"
	PublishedStateODataType            = "#microsoft.graph.publishedState"
	RecordingInfoODataType             = "#microsoft.graph.recordingInfo"
	ResultInfoODataType                = "#microsoft.graph.resultInfo"
	ServiceHostedMediaConfigODataType  = "#microsoft.graph.serviceHostedMediaConfig"
	TokenMeetingInfoODataType          = "#microsoft.graph.tokenMeetingInfo"
	ToneInfoODataType                  = "#microsoft.graph.toneInfo"
)

// GetType returns an enum ResourceType from an ODataType name
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
		case CallRouteODataType:
			return CallRouteResourceType
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
	case 29:
		switch string(ODataType) {
		case PublishedStateODataType:
			return PublishedStateResourceType
		}
	case 30:
		switch string(ODataType) {
		case RecordingInfoODataType:
			return RecordingInfoResourceType
		}
	case 31:
		switch string(ODataType) {
		case CallMediaStateODataType:
			return CallMediaStateResourceType
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
	case 42:
		switch string(ODataType) {
		case InvitationParticipantInfoODataType:
			return InvitationParticipantInfoResourceType
		}
	}
	return UnknownResourceType
}

// NewResource returns a new Resource from an ODataType name. The object will
// match the type specified by the ODataType
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
		case CallRouteODataType:
			return &CallRoute{}
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
	case 29:
		switch string(ODataType) {
		case PublishedStateODataType:
			return &PublishedState{}
		}
	case 30:
		switch string(ODataType) {
		case RecordingInfoODataType:
			return &RecordingInfo{}
		}
	case 31:
		switch string(ODataType) {
		case CallMediaStateODataType:
			return &CallMediaState{}
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
	case 42:
		switch string(ODataType) {
		case InvitationParticipantInfoODataType:
			return &InvitationParticipantInfo{}
		}
	}
	return nil
}
