package resource

// microsoft.graph.identitySet
// https://learn.microsoft.com/en-us/graph/api/resources/identityset?view=graph-rest-1.0
type IdentitySet struct {
	ODataType   string    `json:"@odata.type"`
	AcsUser     *Identity `json:"acsUser,omitempty"`
	Application *Identity `json:"application,omitempty"`
	Phone       *Identity `json:"phone,omitempty"`
	User        *Identity `json:"user,omitempty"`
}

func (r *IdentitySet) GetType() (ResourceType, bool) {
	return IdentitySetResourceType, false
}

func (r *IdentitySet) Validate() bool {
	if r.ODataType != IdentitySetODataType {
		return false
	}
	if r.AcsUser != nil && !r.AcsUser.Validate() {
		return false
	}
	if r.Application != nil && !r.Application.Validate() {
		return false
	}
	if r.Phone != nil && !r.Phone.Validate() {
		return false
	}
	if r.User != nil && !r.User.Validate() {
		return false
	}
	return true
}
