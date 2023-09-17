package resource

import "github.com/Soreing/parsley"

// Resource describes any microsoft graph resource. Each resource must implement
// how to unmarshal from a json, how to return its ResourceType and define how
// to Validate its integrity.
type Resource interface {
	parsley.Unmarshaller
	GetType() (rtype ResourceType, isSlice bool)
	Validate() (ok bool)
}
