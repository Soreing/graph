package resource

// microsoft.graph.identity
// https://learn.microsoft.com/en-us/graph/api/resources/identity?view=graph-rest-1.0
type Identity struct {
	ODataType        string `json:"@odata.type"`
	AcsResourceId    string `json:"acsResourceId,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	Id               string `json:"id,omitempty"`
	IdentityProvider string `json:"identityProvider,omitempty"`
	TenantId         string `json:"tenantId,omitempty"`
}

func (r *Identity) GetType() (ResourceType, bool) {
	return IdentityResourceType, false
}

func (r *Identity) Validate() bool {
	return r.ODataType == IdentityODataType
}
