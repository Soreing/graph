package resource

// microsoft.graph.playPromptOperation
// https://learn.microsoft.com/en-us/graph/api/resources/playpromptoperation?view=graph-rest-1.0
type PlayPromptOperation struct {
	ODataType        string        `json:"@odata.type"`
	CompletionReason string        `json:"completionReason,omitempty"`
	Id               string        `json:"id,omitempty"`
	Prompts          []MediaPrompt `json:"prompts,omitempty"`
	ResultInfo       *ResultInfo   `json:"resultInfo,omitempty"`
	Status           string        `json:"status,omitempty"`
}

func (r *PlayPromptOperation) GetType() (ResourceType, bool) {
	return PlayPromptOperationResourceType, false
}

func (r *PlayPromptOperation) Validate() bool {
	if r.ODataType != PlayPromptOperationODataType {
		return false
	}
	if r.ResultInfo != nil && !r.ResultInfo.Validate() {
		return false
	}
	for _, rs := range r.Prompts {
		if !rs.Validate() {
			return false
		}
	}
	return true
}
