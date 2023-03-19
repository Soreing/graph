package resource

import "github.com/Soreing/fastjson/reader"

type Resource interface {
	GetType() (rtype ResourceType, isSlice bool)
	Validate() (ok bool)
	UnmarshalFastJSON(r *reader.Reader) (err error)
}
