package resource

// microsoft.graph.resultInfo
// https://learn.microsoft.com/en-us/graph/api/resources/resultinfo?view=graph-rest-1.0
type ResultInfo struct {
	ODataType string `json:"@odata.type"`
	Code      int32  `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Subcode   int32  `json:"subcode,omitempty"`
}

func (r *ResultInfo) GetType() (ResourceType, bool) {
	return ResultInfoResourceType, false
}

func (r *ResultInfo) Validate() bool {
	return r.ODataType == ResultInfoODataType
}
