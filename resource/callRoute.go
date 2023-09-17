package resource

// microsoft.graph.callRoute
// https://learn.microsoft.com/en-us/graph/api/resources/callRoute?view=graph-rest-1.0
type CallRoute struct {
	ODataType   string       `json:"@odata.type"`
	Final       *IdentitySet `json:"final"`
	Original    *IdentitySet `json:"original"`
	RoutingType string       `json:"routingType"`
}

func (r *CallRoute) GetType() (ResourceType, bool) {
	return CallRouteResourceType, false
}

func (r *CallRoute) Validate() bool {
	return r.ODataType == CallRouteODataType
}
